import './Welcome.css'

export default function Welcome(props: {
  msg : string
}) {
  return <div className="greetings">
    <h1 className="green">{props.msg}</h1>
    <h3>
    You've successfully created a project with
    <a href="https://github.com/airconduct/go-probot/" target="_blank" rel="noopener"> Go-Probot</a>.
    </h3>
  </div>
}