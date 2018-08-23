import React, { Component } from 'react';
import './Analytics.css';

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
    fetch("http://localhost:8080/glaza/analytics?pagesize=3&sort=-date")
      .then(res => res.json())
      .then(
        (result) => {
          this.setState({
            isLoaded: true,
            items: result._embedded
          });
          // console.log(result._embedded[0].json.totalsForAllResults);
          // console.log(result._embedded[0].json.totalsForAllResults["ga:users"]);
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
          <div class="analytics">
            <b>{item.name} {item.value}</b>
            <table>
              <tr>
                <td>Date</td>
                <td>Users</td>
                <td>Impressions</td>
                <td>AD Clicks</td>                
                <td>Organic</td>                
              </tr>
              {item.json.rows.map(row => (
                <tr>
                  <td>{row[0]}</td>
                  <td>{row[1]}</td>
                  <td>{row[2]}</td>
                  <td>{row[3]}</td>
                  <td>{row[4]}</td>
                </tr>
              ))}
              <tr>
                <td>...</td>
                <td>{item.json.totalsForAllResults["ga:users"]}</td>
                <td>{item.json.totalsForAllResults["ga:impressions"]}</td>                
                <td>{item.json.totalsForAllResults["ga:adClicks"]}</td>                
                <td>{item.json.totalsForAllResults["ga:organicSearches"]}</td>
              </tr>
            </table>
          </div>
        ))}
      </div>
    );   
  }
}

export default WebmastersClicks;