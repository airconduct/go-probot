import { useState } from 'react'
import probotLog from './assets/robot.svg'
import './App.css'

import { Routes, Route, Outlet, Link } from "react-router-dom";

import {  Menu  } from "antd";
import {
  DashboardOutlined,
  CopyrightOutlined
} from '@ant-design/icons';
import 'antd/dist/reset.css';

import Welcome from './Welcome'
import ErrorPage from './error-page'
import Dashboard from './Dashboard'

export default function App() {
  return <div id='app'>
      <header>
        <img src={probotLog} className="logo" alt="Probot logo" width="125" height="125"/>
        <div className="wrapper">
          <Welcome msg='Welcome!'></Welcome>
          <Menu
            className='nav'
            mode="horizontal"
            onClick={(e)=>{console.log(e)}}
            items={[
              {key:"1",label:(<Link to="/dashboard">Dashboard</Link>), icon: <DashboardOutlined />},
              {key:"2",label:(<Link to="/dashboard/about">About</Link>), icon: <CopyrightOutlined />},
            ]}
          />
        </div>
      </header>
      <main>
        <Routes>
          <Route path="/">
            <Route path="dashboard/about" element={<About />} errorElement={<ErrorPage/>} />
            <Route path="dashboard" element={<Dashboard />} errorElement={<ErrorPage/>} />
            {/* Using path="*"" means "match anything", so this route
                  acts like a catch-all for URLs that we don't have explicit
                  routes for. */}
            <Route path="*" element={<NoMatch />} errorElement={<ErrorPage/>} />
          </Route>
        </Routes>
        {/* An <Outlet> renders whatever child route is currently active,
            so you can think about this <Outlet> as a placeholder for
            the child routes we defined above. */}
        <Outlet />
      </main>
  </div>
}

function About() {
  return <div>
    <h1>About</h1>
    <p>Do you like this project? Give us a start on <a target={"_blank"} href='https://github.com/airconduct/go-probot'>GitHub</a>!</p>
  </div>
}

function NoMatch() {
  return (
    // <ErrorPage></ErrorPage>
    <div>
      <h2>Nothing to see here!</h2>
      <p>
        <Link to="/">Go to the home page</Link>
      </p>
    </div>
  );
}

