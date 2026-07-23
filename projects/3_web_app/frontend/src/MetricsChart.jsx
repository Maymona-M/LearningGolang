import { LineChart, Line, XAxis, YAxis, Tooltip, Legend, CartesianGrid, ResponsiveContainer } from 'recharts'

function MetricsChart({ readings }) {
    const chartData = []
    readings.forEach(reading => {
        chartData.push({
            time: new Date(reading.timestamp * 1000).toLocaleTimeString(),
            CPU: reading.cpu,
            Memory: reading.mem,
            Disk: reading.disk
        })
    })

    return (
        <div>
            {/* Individual charts, one per metric */}
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

            {/* Combined overview chart */}
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