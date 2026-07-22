import { useState } from 'react'

const ROWS_PER_PAGE = 10

function DataTable({ columns, data }) {
    const [sortKey, setSortKey] = useState(null)
    const [sortAsc, setSortAsc] = useState(true)
    const [currentPage, setCurrentPage] = useState(1)

    const handleSort = (key) => {
        if (sortKey === key) {
            setSortAsc(!sortAsc)
        } else {
            setSortKey(key)
            setSortAsc(true)
        }
        setCurrentPage(1) // reset to page 1 whenever sort changes
    }

    let sortedData = [...data]
    if (sortKey) {
        sortedData.sort((a, b) => {
            if (a[sortKey] < b[sortKey]) return sortAsc ? -1 : 1
            if (a[sortKey] > b[sortKey]) return sortAsc ? 1 : -1
            return 0
        })
    }

    // figure out which slice of data belongs to the current page
    const totalPages = Math.ceil(sortedData.length / ROWS_PER_PAGE)
    const startIndex = (currentPage - 1) * ROWS_PER_PAGE
    const pageData = sortedData.slice(startIndex, startIndex + ROWS_PER_PAGE)

    const goToPreviousPage = () => {
        if (currentPage > 1) setCurrentPage(currentPage - 1)
    }

    const goToNextPage = () => {
        if (currentPage < totalPages) setCurrentPage(currentPage + 1)
    }

    // build header cells
    const headerCells = []
    columns.forEach(col => {
        headerCells.push(
            <th key={col.key} onClick={() => handleSort(col.key)} style={{ cursor: 'pointer' }}>
                {col.label} {sortKey === col.key ? (sortAsc ? '▲' : '▼') : ''}
            </th>
        )
    })

    // build rows for only the current page
    const bodyRows = []
    pageData.forEach(row => {
        const cells = []
        columns.forEach(col => {
            cells.push(
                <td key={col.key}>{col.render ? col.render(row) : row[col.key]}</td>
            )
        })
        bodyRows.push(<tr key={row.id}>{cells}</tr>)
    })

    return (
        <div>
            <table>
                <thead>
                    <tr>{headerCells}</tr>
                </thead>
                <tbody>{bodyRows}</tbody>
            </table>

            <div className="page-nav">
                <button onClick={goToPreviousPage} disabled={currentPage === 1}>
                    ← Previous
                </button>
                <span> Page {currentPage} of {totalPages} </span>
                <button onClick={goToNextPage} disabled={currentPage === totalPages}>
                    Next →
                </button>
            </div>
        </div>
    )
}

export default DataTable