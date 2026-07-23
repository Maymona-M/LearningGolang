export function aggregateReadings(readings, granularity) {
    const groups = new Map()

    readings.forEach(reading => {
        const date = new Date(reading.timestamp * 1000)

        let key
        if (granularity === 'hour') {
            key = date.toLocaleString('en-US', {
                month: 'short', day: 'numeric', year: 'numeric', hour: 'numeric'
            })
        } else if (granularity === 'day') {
            key = date.toLocaleDateString('en-US', {
                month: 'short', day: 'numeric', year: 'numeric'
            })
        }

        if (!groups.has(key)) {
            groups.set(key, [])
        }
        groups.get(key).push(reading)
    })

    const result = []
    groups.forEach((groupReadings, key) => {
        let cpuTotal = 0, memTotal = 0, diskTotal = 0

        groupReadings.forEach(r => {
            cpuTotal += r.cpu
            memTotal += r.mem
            diskTotal += r.disk
        })

        const count = groupReadings.length

        // track the earliest timestamp in this group, so we can sort properly later
        const earliestTimestamp = Math.min(...groupReadings.map(r => r.timestamp))

        result.push({
            id: key,
            period: key,
            cpu: cpuTotal / count,
            mem: memTotal / count,
            disk: diskTotal / count,
            count: count,
            sortKey: earliestTimestamp // used only for sorting, not displayed
        })
    })

    // sort oldest -> newest by that real timestamp, regardless of Map insertion order
    result.sort((a, b) => a.sortKey - b.sortKey)

    return result
}