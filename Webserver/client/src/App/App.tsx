import React, { useEffect } from 'react';
import Box from '@mui/material/Box';
import {
  BrowserRouter as Router,
  Routes,
  Route,
} from "react-router-dom";

import Home from '../Pages/Home';
import Index from '../Pages/Index';
import Searches from '../Pages/Searches';
import Login from '../Pages/Login';

import EventBus from '../utils/eventbus';
import { EVENTS } from '../constants';

import { StyledLink, Container } from './styles';

function App() {
  useEffect(() => {
    EventBus.subscribe(EVENTS.NAVIGATE, (path: string) => {
      window.location.assign(path);
    });
  }, []);

  return (
    <Router>
      <div className="App">
        <Box style={{
          background: '#333',
          width: '100%',
          marginBottom: '20px',
          boxShadow: '0 0 10px rgba(0,0,0,.25)'
        }}>
          <Container>
            <StyledLink to="/">Home</StyledLink>
          </Container>
        </Box>
        <Container>
          <Routes>
            <Route path="" element={<Home />} />
            <Route path="/login" element={<Login />} />
            <Route path="/index/:id" element={<Index />} />
            <Route path="/index/:id/search/:search_slug" element={<Searches />} />
          </Routes>
        </Container>
      </div>
    </Router>
  );
}

export default App;
