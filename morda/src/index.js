import React from 'react';
import ReactDOM from 'react-dom';
import './index.css';
import Docs from './components/docs/Docs';
import PageInsights from './components/pageinsights/PageInsights';
import registerServiceWorker from './registerServiceWorker';

ReactDOM.render(<Docs />, document.getElementById('docs'));
ReactDOM.render(<PageInsights />, document.getElementById('pageinsights'));

registerServiceWorker();
