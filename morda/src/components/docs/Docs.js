import React, { Component } from 'react';
import './Docs.css';

class Docs extends Component {
  constructor(props) {
    super(props);
    this.state = {
      error: null,
      isLoaded: false,
      items: []
    };
  }

  componentDidMount() {
    fetch("http://api.glaza:8000/api/glaza/docs?pagesize=6&sort=-date")
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
            {items.map(item => (
              <li key={item.name}>
                {item.name} {item.value} <a href={item.json.url}>link</a>
              </li>
            ))}
          </ul>
        </div>
      );
    }      
  }
}

export default Docs;