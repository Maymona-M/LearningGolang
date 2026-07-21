import { useState } from 'react'
import './App.css'
import IPListPage from './IPListPage.jsx'
import MetricsPage from './MetricsPage.jsx'

function App() {
  const [selectedIp, setSelectedIp] = useState(null) // null = show list page

  if (selectedIp === null) {
    return <IPListPage onSelectIp={setSelectedIp} />
  }

  return <MetricsPage ip={selectedIp} onBack={() => setSelectedIp(null)} />
}

export default App