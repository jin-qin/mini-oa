import './App.css';
import './pages/page.css';
import { Navbar, Nav } from 'react-bootstrap';
import { 
  HashRouter as Router,
  Route,
  NavLink
} from 'react-router-dom'
import { CSSTransition } from 'react-transition-group';

import Dashboard from './pages/Dashboard';
import Login from './pages/Login';
import Register from './pages/Register';

function App() {
  const routes = [
    { path: '/', name: 'Dashboard', Component: Dashboard },
    { path: '/dashboard', name: 'Dashboard', Component: Dashboard },
    { path: '/login', name: 'Login', Component: Login },
    { path: '/register', name: 'Register', Component: Register },
  ]

  return (
    <div className="App">
      <Router>
        <>
        <Navbar bg="light">
          <Nav className="mx-auto">
            {
              routes.map(route => (
                <Nav.Link
                  key={route.path}
                  as={NavLink}
                  to={route.path}
                  activeClassName="active"
                  exact
                >
                  {route.name}
                </Nav.Link>
              ))
            }
          </Nav>
        </Navbar>

        <div>
        {
          routes.map(({ path, Component }) => (
            <Route key={path} exact path={path}>
            {
              ({ match }) => (
                <CSSTransition
                  in={match != null}
                  timeout={300}
                  classNames='page'
                  unmountOnExit
                >
                  <div className='page'>
                    <Component />
                  </div>
                </CSSTransition>
              )
            }
            </Route>
          ))
        }
        </div>
        </>
      </Router>
      
    </div>
  );
}

export default App;
