import { useState, useEffect } from 'react'

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

    // build table rows
    const rows = []
    readings.forEach(reading => {
        // Convert Unix timestamp (seconds) to JavaScript timestamp (milliseconds)
        const date = new Date(reading.timestamp * 1000)
        // Format options for Date and 12-hour AM/PM Time
        const readableTime = date.toLocaleString('en-US', {
            year: 'numeric',
            month: 'short',
            day: 'numeric',
            hour: '2-digit',
            minute: '2-digit',
            second: '2-digit',
            hour12: true
        })

        rows.push(
            <tr key={reading.id}>
                <td>{reading.cpu.toFixed(1)}%</td>
                <td>{reading.mem.toFixed(1)}%</td>
                <td>{reading.disk.toFixed(1)}%</td>
                <td>{readableTime}</td>
            </tr>
        )
    })

    return (
        <div>
            <button onClick={onBack}>← Back</button>
            <h1>Metrics for {ip}</h1>
            <table>
                <thead>
                    <tr>
                        <th>CPU</th>
                        <th>Mem</th>
                        <th>Disk</th>
                        <th>Time</th>
                    </tr>
                </thead>
                <tbody>{rows}</tbody>
            </table>
        </div>
    )
}

export default MetricsPage