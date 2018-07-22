import React, { Component } from 'react';

class WebmastersQueries extends Component {
  constructor(props) {
    super(props);
    this.state = {
      error: null,
      isLoaded: false,
      items: []
    };
  }

  componentDidMount() {
    fetch("http://localhost:8080/glaza/webmasters?pagesize=1&sort=-date&filter=%7B%22value%22%3A%20%22queries%22%7D")
      .then(res => res.json())
      .then(
        (result) => {
          this.setState({
            isLoaded: true,
            items: result._embedded
          });
        },
        (error) => {
          this.setState({
            isLoaded: true,
            error
          });
        }
      )
  }

  render() {
    const { error, isLoaded, items } = this.state;  
    return (
      <div>
        {error && <div>Error: {error.message}</div>}
        {items.map(item => (
          <div>
            <b>{item.name} {item.value}</b>
            <table>
              {item.json.rows && item.json.rows.map(row => (
                <tr>
                  <td>{row.clicks}</td>
                  <td>{row.ctr}</td>
                  <td>{row.impressions}</td>
                  <td>{row.position}</td>
                  <td>{row.keys[0]}</td>
                </tr>
              ))}
            </table>
          </div>
        ))}
      </div>
    );   
  }
}

export default WebmastersQueries;