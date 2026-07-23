import { LineChart, Line, XAxis, YAxis, Tooltip, Legend, CartesianGrid, ResponsiveContainer } from 'recharts'
import { aggregateReadings } from './aggregate.js'

// decide what level of detail the chart itself should show,
// one step finer than whatever the table is currently grouped by
function getChartGranularity(tableGranularity) {
    if (tableGranularity === 'day') return 'hour'
    if (tableGranularity === 'hour') return 'raw'
    return 'raw'
}

function MetricsChart({ readings, tableGranularity }) {
    const chartGranularity = getChartGranularity(tableGranularity)
    // the overlay average matches whatever the table is currently set to
    const overlayGranularity = tableGranularity === 'raw' ? null : tableGranularity

    // build the main solid-line data, oldest -> newest
    let mainData
    if (chartGranularity === 'raw') {
        mainData = [...readings].reverse().map(r => ({
            time: new Date(r.timestamp * 1000).toLocaleString('en-US', {
                month: 'short', day: 'numeric', hour: '2-digit', minute: '2-digit'
            }),
            tsForMatch: r.timestamp,
            CPU: r.cpu,
            Memory: r.mem,
            Disk: r.disk
        }))
    } else {
        mainData = aggregateReadings(readings, chartGranularity).map(r => ({
            time: r.period,
            tsForMatch: r.sortKey,
            CPU: r.cpu,
            Memory: r.mem,
            Disk: r.disk
        }))
    }

    // if there's a coarser average to show, compute it and attach to each main point
    let chartData = mainData
    if (overlayGranularity) {
        const overlay = aggregateReadings(readings, overlayGranularity)

        chartData = mainData.map(point => {
            const match = overlay.find(avg => avg.sortKey <= point.tsForMatch)
            return {
                ...point,
                CPU_avg: match ? match.cpu : null,
                Memory_avg: match ? match.mem : null,
                Disk_avg: match ? match.disk : null
            }
        })
    }

    // reusable mini chart for one metric, with its solid line + optional grey dashed average
    function renderMetricChart(label, dataKey, avgKey, color) {
        return (
            <div className="chart-box">
                <h3>{label}</h3>
                <ResponsiveContainer width="100%" height={200}>
                    <LineChart data={chartData}>
                        <CartesianGrid strokeDasharray="3 3" />
                        <XAxis dataKey="time" tick={{ fontSize: 9 }} />
                        <YAxis />
                        <Tooltip />
                        <Line type="monotone" dataKey={dataKey} stroke={color} dot={false} />
                        {overlayGranularity && (
                            <Line type="monotone" dataKey={avgKey} stroke="#4f4d4d" strokeDasharray="5 5" dot={false} />
                        )}
                    </LineChart>
                </ResponsiveContainer>
            </div>
        )
    }

    return (
        <div>
            <div className="chart-grid">
                {renderMetricChart('CPU %', 'CPU', 'CPU_avg', '#e74c3c')}
                {renderMetricChart('Memory %', 'Memory', 'Memory_avg', '#3498db')}
                {renderMetricChart('Disk %', 'Disk', 'Disk_avg', '#2ecc71')}
            </div>

            <h3 style={{ marginTop: '30px' }}>Combined Overview</h3>
            <ResponsiveContainer width="100%" height={300}>
                <LineChart data={chartData}>
                    <CartesianGrid strokeDasharray="3 3" />
                    <XAxis dataKey="time" tick={{ fontSize: 10 }} />
                    <YAxis />
                    <Tooltip />
                    <Legend />
                    <Line type="monotone" dataKey="CPU" stroke="#e74c3c" dot={false} />
                    <Line type="monotone" dataKey="Memory" stroke="#3498db" dot={false} />
                    <Line type="monotone" dataKey="Disk" stroke="#2ecc71" dot={false} />
                </LineChart>
            </ResponsiveContainer>
        </div>
    )
}

export default MetricsChart