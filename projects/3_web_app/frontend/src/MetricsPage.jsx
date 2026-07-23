import { useState, useEffect } from 'react'
import DataTable from './DataTable.jsx'
import MetricsChart from './MetricsChart.jsx'


// ip = which machine to show; onBack = return to list page
function MetricsPage({ ip, onBack }) {
    const [readings, setReadings] = useState([])

    // re-fetch whenever ip changes
    useEffect(() => {
        fetch(`http://localhost:8081/api/readings?ip=${ip}`)
            .then(response => response.json())
            .then(data => setReadings(data))
            .catch(error => console.error('Error fetching readings:', error))
    }, [ip])

    // build table columns
    const columns = [
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

    return (
        <div className="page">
            <button className="back-button" onClick={onBack}>← Back</button>
            <h1>Metrics for</h1>
            <h1> {ip} </h1>

            <h2>Trend</h2>
            <MetricsChart readings={readings} />

            <h2>Recent Readings</h2>
            <DataTable columns={columns} data={readings} />
        </div>
    )
}

export default MetricsPage