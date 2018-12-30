import React, { Component } from 'react';

class WebmastersClicks extends Component {
  constructor(props) {
    super(props);
    this.state = {
      error: null,
      isLoaded: false,
      items: []
    };
  }

  componentDidMount() {
    fetch("/api/glaza/webmasters?pagesize=1&sort=-date&filter=%7B%22value%22%3A%20%22clicks%22%7D")
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
        {items && items.map(item => (
          <div>
            <b>{item.name} {item.value}</b>
            <table>
              <tr>
                <td>Clicks</td>
                <td>CTR</td>
                <td>Impressions</td>
                <td>Position</td>
                <td>Date</td>
              </tr>
            {item.json.rows.slice(0).reverse().map(row => (
              <tr>
                <td>{row.clicks}</td>
                <td>{Number(row.ctr.toFixed(2))}</td>
                <td>{row.impressions}</td>
                <td>{Number(row.position.toFixed(2))}</td>
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

export default WebmastersClicks;