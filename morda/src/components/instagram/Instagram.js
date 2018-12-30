import React, { Component } from 'react';
import './Instagram.css';

class Instagram extends Component {
  constructor(props) {
    super(props);
    this.state = {
      error: null,
      isLoaded: false,
      items: []
    };
  }

  componentDidMount() {
    fetch("/api/glaza/instagram?pagesize=5&sort=-date")
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
              <li>
                <p>{item.username}</p>
                <a target="blank" href={"https://instagram.com/" + item.username}><img width="150" src={item.media_urls[0]}/></a>
              </li>
            ))}
          </ul>
        </div>
      );
    }      
  }
}

export default Instagram;