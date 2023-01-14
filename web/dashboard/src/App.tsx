import { useState } from 'react'
import probotLog from './assets/robot.svg'
import './App.css'

import { Routes, Route, Outlet, Link } from "react-router-dom";

import {  Menu, Card, Collapse, Space  } from "antd";
import {
  DashboardOutlined, PullRequestOutlined, InfoCircleOutlined,
  CopyrightOutlined
} from '@ant-design/icons';
import 'antd/dist/reset.css';

const { Panel } = Collapse;
import Welcome from './Welcome'
import ErrorPage from './error-page'

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
    <p>Do you like this project? Give us a start on <a target={"_blank"} href='https://github.com/airconduct/go-probot'>Github</a>!</p>
  </div>
}

function Dashboard() {
  const onChange = (key: string | string[]) => {
    console.log(key);
  };
  return <div>
    <h1>Event Listener</h1>
    <Collapse onChange={onChange} accordion>
      <Panel key="2" header={
        <Space><InfoCircleOutlined /><div>issue</div></Space>
      }>
        FOO
      </Panel>
      <Panel key="1" header={
        <Space><PullRequestOutlined /><div>pulls</div></Space>
      }>
        BAR
      </Panel>
    </Collapse>
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

