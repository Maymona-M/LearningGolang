import { useState, useEffect } from 'react'

// onSelectIp is passed down from App.jsx — called when an IP is clicked
function IPListPage({ onSelectIp }) {
    const [ips, setIps] = useState([]) // list of unique IPs from backend

    // fetch once on load
    useEffect(() => {
        fetch('http://localhost:8081/api/ips')
            .then(response => response.json())
            .then(data => setIps(data))
            .catch(error => console.error('Error fetching IPs:', error))
    }, [])

    // build list items manually
    const listItems = []
    ips.forEach(ip => {
        listItems.push(
            <li key={ip}>
                <button onClick={() => onSelectIp(ip)}>{ip}</button>
            </li>
        )
    })

    return (
        <div>
            <h1>Known Connections</h1>
            <ul>{listItems}</ul>
        </div>
    )
}

export default IPListPage