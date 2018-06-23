import React from 'react';
import ReactDOM from 'react-dom';
import './index.css';
import Docs from './components/docs/Docs';
import PageInsights from './components/pageinsights/PageInsights';
import WebmastersClicks from './components/webmasters/WebmastersClicks';
import WebmastersQueries from './components/webmasters/WebmastersQueries';
import registerServiceWorker from './registerServiceWorker';

ReactDOM.render(<Docs />, document.getElementById('docs'));
ReactDOM.render(<PageInsights />, document.getElementById('pageinsights'));
ReactDOM.render(<WebmastersClicks />, document.getElementById('webmasters-clicks'));
ReactDOM.render(<WebmastersQueries />, document.getElementById('webmasters-queries'));

registerServiceWorker();
