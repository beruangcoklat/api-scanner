import React, { useState } from 'react';
import { Form, Button } from 'react-bootstrap'

export default function InsertAPIPage(props) {

  const [name, setName] = useState('')
  const [dbms, setDbms] = useState('')
  const [data, setData] = useState('')

  const onSubmit = function (e) {
    e.preventDefault()

    const payload = {
      name: name,
      data: data,
      dbms: dbms,
    }

    fetch('/api/api-data', {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json'
      },
      body: JSON.stringify(payload)
    })
      .then(() => {
        setName('')
        setDbms('')
        setData('')
      })
  }

  return (
    <Form required onSubmit={onSubmit}>
      <Form.Group className="mb-3">
        <Form.Label>API Name</Form.Label>
        <Form.Control
          type="text"
          placeholder="Enter API Name"
          required
          value={name}
          onChange={e => setName(e.target.value)}
        />
      </Form.Group>

      <Form.Group className="mb-3">
        <Form.Label>DBMS</Form.Label>
        <Form.Control
          type="text"
          placeholder="Enter DBMS"
          required
          value={dbms}
          onChange={e => setDbms(e.target.value)}
        />
      </Form.Group>

      <Form.Group className="mb-3">
        <Form.Label>HTTP Data</Form.Label>
        <Form.Control
          as="textarea" rows={3}
          placeholder="GET / HTTP/1.1
Host: www.google.com"
          required
          value={data}
          onChange={e => setData(e.target.value)}
        />
      </Form.Group>

      <Button variant="primary" type="submit">
        Submit
      </Button>
    </Form>
  )
}
