import { LineChart, Line, XAxis, YAxis, Tooltip, Legend, CartesianGrid, ResponsiveContainer } from 'recharts'
import { aggregateReadings } from './aggregate.js'

function MetricsChart({ readings, granularity }) {
    // pick raw or aggregated data depending on the selected view
    let dataToChart = readings
    if (granularity === 'hour' || granularity === 'day') {
        dataToChart = aggregateReadings(readings, granularity)
    }


    // reverse so the chart reads oldest -> newest, left to right
    const chronological = granularity === 'raw' ? [...readings].reverse() : aggregateReadings(readings, granularity)
    const chartData = []
    chronological.forEach(item => {
        if (granularity === 'raw') {
            chartData.push({
                time: new Date(item.timestamp * 1000).toLocaleString('en-US', {
                    month: 'short', day: 'numeric', hour: '2-digit', minute: '2-digit'
                }),
                CPU: item.cpu,
                Memory: item.mem,
                Disk: item.disk
            })
        } else {
            // aggregated items already have a readable "period" label
            chartData.push({
                time: item.period,
                CPU: item.cpu,
                Memory: item.mem,
                Disk: item.disk
            })
        }
    })

    return (
        <div>
            <div className="chart-grid">
                <div className="chart-box">
                    <h3>CPU %</h3>
                    <ResponsiveContainer width="100%" height={200}>
                        <LineChart data={chartData}>
                            <CartesianGrid strokeDasharray="3 3" />
                            <XAxis dataKey="time" tick={{ fontSize: 9 }} />
                            <YAxis />
                            <Tooltip />
                            <Line type="monotone" dataKey="CPU" stroke="#e74c3c" dot={false} />
                        </LineChart>
                    </ResponsiveContainer>
                </div>

                <div className="chart-box">
                    <h3>Memory %</h3>
                    <ResponsiveContainer width="100%" height={200}>
                        <LineChart data={chartData}>
                            <CartesianGrid strokeDasharray="3 3" />
                            <XAxis dataKey="time" tick={{ fontSize: 9 }} />
                            <YAxis />
                            <Tooltip />
                            <Line type="monotone" dataKey="Memory" stroke="#3498db" dot={false} />
                        </LineChart>
                    </ResponsiveContainer>
                </div>

                <div className="chart-box">
                    <h3>Disk %</h3>
                    <ResponsiveContainer width="100%" height={200}>
                        <LineChart data={chartData}>
                            <CartesianGrid strokeDasharray="3 3" />
                            <XAxis dataKey="time" tick={{ fontSize: 9 }} />
                            <YAxis />
                            <Tooltip />
                            <Line type="monotone" dataKey="Disk" stroke="#2ecc71" dot={false} />
                        </LineChart>
                    </ResponsiveContainer>
                </div>
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