import React, { Component } from 'react';

class Bots extends Component {
  constructor(props) {
    super(props);
    this.state = {
      error: null,
      isLoaded: false,
      items: []
    };
  }

  componentDidMount() {
    fetch("/api/glaza/bots?pagesize=1&sort=-date")
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
    if (error) {
      return <div>Error: {error.message}</div>;
    } else if (!isLoaded) {
      return <div>Loading...</div>;
    } else {    
      return (
        <div>
          <ul>
            {items && items.map(item => (
              <li >
                Bot: {item.bot_name}, Status: {item.status}, Date: {(new Date(item.date.$date)).toString()}
              </li>
            ))}
          </ul>
        </div>
      );
    }      
  }
}

export default Bots;