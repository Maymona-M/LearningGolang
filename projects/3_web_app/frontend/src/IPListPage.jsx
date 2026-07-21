import { useState, useEffect } from 'react'

function IPListPage() {
    const [ips, setIps] = useState([])

    useEffect(() => {
        fetch('http://localhost:8081/api/ips')
            .then(response => response.json())
            .then(data => setIps(data))
            .catch(error => console.error('Error fetching IPs:', error))
    }, [])

    const listItems = []
    ips.forEach(ip => {
        listItems.push(<li key={ip}>{ip}</li>)
    })

    return (
        <div>
            <h1>Known Machines</h1>
            <ul>{listItems}</ul>
        </div>
    )
}

export default IPListPage