import React, { useState, useEffect } from 'react';
import { ListGroup } from 'react-bootstrap'
import { useHistory } from 'react-router-dom';

export default function ListAPIPage(props) {

  const [list, setList] = useState([])
  const history = useHistory();

  useEffect(() => {
    fetch('/api/api-data', {
      method: 'GET',
      headers: {
        'Content-Type': 'application/json'
      },
    })
      .then(r => r.json())
      .then(r => setList(r))
  }, [])

  return (
    <ListGroup>
      {
        list.map(i => {
          return (
            <ListGroup.Item
              onClick={() => history.push(`/detail/${i.id}`)}
              key={i.id}
            >
              {i.name}
            </ListGroup.Item>
          )
        })
      }
    </ListGroup>
  )
}
