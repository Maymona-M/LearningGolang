import React, { useState } from 'react';

function DataTable({ columns, data }) {
    const [sortKey, setSortKey] = useState(null)
    const [sortAsc, setSortAsc] = useState(true)

    const handleSort = (key) => {
        if (sortKey === key) {
            setSortAsc(!sortAsc)
        } else {
            setSortKey(key)
            setSortAsc(true)
        }
    }

    let sortedData = [...data]
    if (sortKey) {
        sortedData.sort((a, b) => {
            if (a[sortKey] < b[sortKey]) return sortAsc ? -1 : 1
            if (a[sortKey] > b[sortKey]) return sortAsc ? 1 : -1
            return 0
        })
    }

    // build header cells manually
    const headerCells = []
    columns.forEach(col => {
        headerCells.push(
            <th key={col.key} onClick={() => handleSort(col.key)} style={{ cursor: 'pointer' }}>
                {col.label} {sortKey === col.key ? (sortAsc ? '▲' : '▼') : ''}
            </th>
        )
    })

    // build table rows manually
    const bodyRows = []
    sortedData.forEach(row => {
        const cells = []
        columns.forEach(col => {
            cells.push(
                <td key={col.key}>{col.render ? col.render(row) : row[col.key]}</td>
            )
        })
        bodyRows.push(<tr key={row.id}>{cells}</tr>)
    })

    return (
        <table style={{ width: '100%', borderCollapse: 'collapse', textAlign: 'left' }}>
            <thead>
                <tr>{headerCells}</tr>
            </thead>
            <tbody>{bodyRows}</tbody>
        </table>
    )
}

export default DataTable