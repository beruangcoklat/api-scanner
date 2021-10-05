import React, { useState, useEffect } from 'react';
import { useParams } from 'react-router-dom';
import { Button } from 'react-bootstrap'

export default function DetailAPIPage(props) {

  const [data, setData] = useState({})
  const params = useParams();
  const [activeLog, setActiveLog] = useState(-1)

  useEffect(() => {
    fetch(`/api/api-data/${params.id}`, {
      method: 'GET',
      headers: {
        'Content-Type': 'application/json'
      },
    })
      .then(r => r.json())
      .then(r => setData(r))
  }, [])

  const scan = () => {
    fetch(`/api/api-data/${params.id}/scan`, {
      method: 'GET',
      headers: {
        'Content-Type': 'application/json'
      },
    })
  }

  const updateActiveLog = (key) => {
    if (key == activeLog) {
      setActiveLog(-1)
      return
    }
    setActiveLog(key)
  }

  return (
    <>
      <table cellPadding='10px'>
        <tbody>
          <tr>
            <td>API Name</td>
            <td>{data.name}</td>
          </tr>
          <tr>
            <td>DBMS</td>
            <td>{data.dbms}</td>
          </tr>
          <tr>
            <td>Data</td>
            <td>
              <pre>
                {data.data}
              </pre>
            </td>
          </tr>
          <tr>
            <td colSpan={2}>
              <Button variant="primary" type="submit" onClick={scan}>
                Scan
              </Button>
            </td>
          </tr>
        </tbody>
      </table>

      <br />
      <h5>Scan History</h5>

      {
        data.scan_result?.map((i, key) => {
          return (
            <div key={key}>
              <label style={{ marginRight: '20px' }}>{i.created_at}</label>
              <label style={{ marginRight: '20px', marginBottom: '20px' }}>
                {i.is_vulnerable ? 'Vulnerability Detected' : 'Safe'}
              </label>

              <Button variant="primary" type="submit" onClick={() => updateActiveLog(key)}>
                Scan
              </Button>

              {activeLog == key && <pre style={{ marginBottom: '100px' }} >{i.log}</pre>}
            </div>
          )
        })
      }
    </>
  )
}
