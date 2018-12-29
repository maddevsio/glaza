import React from 'react';
import ReactDOM from 'react-dom';
import './index.css';
import Docs from './components/docs/Docs';
import Bots from './components/bots/Bots';
import PageInsights from './components/pageinsights/PageInsights';
import WebmastersClicks from './components/webmasters/WebmastersClicks';
import WebmastersQueries from './components/webmasters/WebmastersQueries';
import AnalyticsTotal from './components/analytics/Total';
import Instagram from './components/instagram/Instagram';
import registerServiceWorker from './registerServiceWorker';

ReactDOM.render(<Docs />, document.getElementById('docs'));
ReactDOM.render(<Bots />, document.getElementById('bots'));
ReactDOM.render(<PageInsights />, document.getElementById('pageinsights'));
ReactDOM.render(<WebmastersClicks />, document.getElementById('webmasters-clicks'));
ReactDOM.render(<WebmastersQueries />, document.getElementById('webmasters-queries'));
ReactDOM.render(<AnalyticsTotal />, document.getElementById('analytics-total'));
ReactDOM.render(<Instagram />, document.getElementById('instagram'));


registerServiceWorker();
