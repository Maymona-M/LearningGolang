import { useState, useEffect } from 'react'
import serverIcon from './assets/server.png'

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
    const cards = []
    ips.forEach(ip => {
        cards.push(
            <button key={ip} className="server-card" onClick={() => onSelectIp(ip)}>
                <img src={serverIcon} alt="server" />
                <span>{ip}</span>
            </button>
        )
    })

    return (
        <div>
            <h1 style={{ textAlign: 'center' }}>Known Machines</h1>
            <div className="server-list">{cards}</div>
        </div>
    )
}

export default IPListPage