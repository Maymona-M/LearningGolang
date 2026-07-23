import { useState, useEffect } from 'react'
import DataTable from './DataTable.jsx'
import MetricsChart from './MetricsChart.jsx'
import { aggregateReadings } from './aggregate.js'


// ip = which machine to show; onBack = return to list page
function MetricsPage({ ip, onBack }) {
    const [readings, setReadings] = useState([])
    const [granularity, setGranularity] = useState('raw') // 'raw' | 'hour' | 'day'


    // re-fetch whenever ip changes and fetch in real time every 3 seconds
    useEffect(() => {
        const fetchReadings = () => {
            fetch(`http://localhost:8081/api/readings?ip=${ip}`)
                .then(response => response.json())
                .then(data => setReadings(data))
                .catch(error => console.error('Error fetching readings:', error))
        }

        fetchReadings() // run once immediately

        const interval = setInterval(fetchReadings, 3000) // then every 5 seconds

        return () => clearInterval(interval) // cleanup when leaving the page
    }, [ip])

    // build table columns
    const rawColumns = [
        { key: 'cpu', label: 'CPU', render: row => `${row.cpu.toFixed(1)}%` },
        { key: 'mem', label: 'Memory', render: row => `${row.mem.toFixed(1)}%` },
        { key: 'disk', label: 'Disk', render: row => `${row.disk.toFixed(1)}%` },
        {
            key: 'timestamp',
            label: 'Time',
            render: row =>
                new Date(row.timestamp * 1000).toLocaleString('en-US', {
                    year: 'numeric',
                    month: 'short',
                    day: 'numeric',
                    hour: '2-digit',
                    minute: '2-digit',
                    second: '2-digit',
                    hour12: true
                })
        }
    ]

    const aggregatedColumns = [
        { key: 'period', label: 'Period' },
        { key: 'cpu', label: 'Average CPU', render: row => `${row.cpu.toFixed(1)}%` },
        { key: 'mem', label: 'Average Memory', render: row => `${row.mem.toFixed(1)}%` },
        { key: 'disk', label: 'Average Disk', render: row => `${row.disk.toFixed(1)}%` },
        { key: 'count', label: 'Readings' }
    ]

    // pick data + columns based on dropdown selection
    let tableData = readings
    let columns = rawColumns
    if (granularity === 'hour' || granularity === 'day') {
        tableData = aggregateReadings(readings, granularity)
        columns = aggregatedColumns
    }

    return (
        <div className="page">
            <button className="back-button" onClick={onBack}>← Back</button>
            <h1>Metrics for {ip}</h1>

            <h2>Trend</h2>
            <MetricsChart readings={readings} granularity={granularity} />

            <h2>Readings</h2>
            <div className="view-selector">
                <label>View: </label>
                <select value={granularity} onChange={e => setGranularity(e.target.value)}>
                    <option value="raw">Raw (every reading)</option>
                    <option value="hour">Hourly average</option>
                    <option value="day">Daily average</option>
                </select>
            </div>

            <DataTable columns={columns} data={tableData} />
        </div>
    )
}

export default MetricsPage