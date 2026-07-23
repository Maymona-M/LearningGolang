// Groups readings by hour or day, and averages cpu/mem/disk for each group
export function aggregateReadings(readings, granularity) {
    const groups = new Map() // key -> array of readings in that bucket

    readings.forEach(reading => {
        const date = new Date(reading.timestamp * 1000)

        let key
        if (granularity === 'hour') {
            // e.g. "Jul 23, 2026, 9 AM"
            key = date.toLocaleString('en-US', {
                month: 'short', day: 'numeric', year: 'numeric', hour: 'numeric'
            })
        } else if (granularity === 'day') {
            // e.g. "Jul 23, 2026"
            key = date.toLocaleDateString('en-US', {
                month: 'short', day: 'numeric', year: 'numeric'
            })
        }

        if (!groups.has(key)) {
            groups.set(key, [])
        }
        groups.get(key).push(reading)
    })

    // now calculate averages for each group
    const result = []
    groups.forEach((groupReadings, key) => {
        let cpuTotal = 0
        let memTotal = 0
        let diskTotal = 0

        groupReadings.forEach(r => {
            cpuTotal += r.cpu
            memTotal += r.mem
            diskTotal += r.disk
        })

        const count = groupReadings.length

        result.push({
            id: key, // DataTable needs a unique "id" per row
            period: key,
            cpu: cpuTotal / count,
            mem: memTotal / count,
            disk: diskTotal / count,
            count: count
        })
    })

    return result
}