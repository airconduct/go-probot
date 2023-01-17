import { useState, useEffect } from 'react'

import {
  PullRequestOutlined, InfoCircleOutlined,
} from '@ant-design/icons';
import {  Menu, Card, Collapse, Space, Spin  } from "antd";
const { Panel } = Collapse;


interface EventList {
  items: Event[]
}

interface Event {
  name:string;
  type:string;
  metrics: {
    count:number;
  };
}

export default function Dashboard() {
  const zeroEventList : EventList = {
    items: [],
  }
  const [eventList, setEventList] = useState(zeroEventList)

  const refresh = async () => {
    const response = await fetch("/api/events")
    const data = await response.json() as EventList
    setEventList((eventList)=>data)
  };

  useEffect(()=>{
    const intervalId = setInterval(() => {
      refresh()
    }, 1000 * 5) // in milliseconds
    return () => clearInterval(intervalId)
  })

  const onChange = async (key: string | string[]) => {
    // TODO: add some actions here
  };

  return <div>
    <h1>Event Listener</h1>
    {eventList.items.length==0?
    <Card hoverable onClick={async ()=>{await refresh()}}>Loading... <Spin/></Card>:
    <Collapse onChange={onChange}>
      {eventList.items.map((event, idx)=>{
        return <Panel key={idx.toString()} header={
          <Space><PullRequestOutlined />{event.name}</Space>
        }>
          <p>Count: {event.metrics.count}</p>
        </Panel>
      })}
    </Collapse>}
  </div>
}

