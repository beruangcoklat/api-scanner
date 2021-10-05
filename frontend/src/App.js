import './App.css';
import {
  Navbar,
  Nav,
  Container
} from 'react-bootstrap'
import {
  BrowserRouter as Router,
  Switch,
  Route,
  Link
} from "react-router-dom";
import InsertAPIPage from './pages/insert_api/index'
import ListAPIPage from './pages/list_api/index'
import DetailAPIPage from './pages/detail_api/index'

function App() {
  return (
    <>
      <Router>
        <Navbar bg="dark" variant="dark">
          <Container>
            <Navbar.Brand>
              API Scanner
            </Navbar.Brand>
            <Nav className="me-auto">
              <Nav.Link to="/insert" as={Link}>Insert API</Nav.Link>
              <Nav.Link to="/list" as={Link}>List API</Nav.Link>
            </Nav>
          </Container>
        </Navbar>

        <div style={{
          width: '50%',
          margin: 'auto',
          marginTop: '20px',
        }}>

          <Switch>
            <Route path="/insert">
              <InsertAPIPage />
            </Route>
            <Route path="/list">
              <ListAPIPage />
            </Route>
            <Route path="/detail/:id">
              <DetailAPIPage />
            </Route>
          </Switch>
        </div>
      </Router>
    </>
  );
}

export default App;
