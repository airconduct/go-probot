import {
  PullRequestOutlined, InfoCircleOutlined,
} from '@ant-design/icons';
import {  Menu, Card, Collapse, Space  } from "antd";
const { Panel } = Collapse;

export default function Dashboard() {
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