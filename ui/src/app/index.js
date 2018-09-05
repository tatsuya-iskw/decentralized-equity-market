import React from 'react';
import {render} from 'react-dom';
import {Provider} from 'react-redux';
import {Router, Route, IndexRoute, hashHistory} from 'react-router';
import {syncHistoryWithStore} from 'react-router-redux';
import MuiThemeProvider from 'material-ui/styles/MuiThemeProvider';

import configureStore from './configureStore';
// import action from './action';
import RootContainer from './component/organism/rootContainer';
import DecentralizedEqutyMarketHome from './component/organism/DecentralizedEqutyMarketHome';

import ShipperHomeDashboard from './component/organism/ShipperHomeDashboard';
import CarrierHomeDashboard from './component/organism/CarrerHomeDashboard';

import Introduction from './component/organism/Introduction';
import HumanAuthentication from './component/organism/HumanAuthentication';
import CorporateDashboard from './component/organism/CorporateDashboard';
import MarketSahreBoard from './component/organism/CarrerHomeDashboard';
import MarketTransactionBoard from './component/organism/CarrerHomeDashboard';
import MarketBitAskBoard from './component/organism/CarrerHomeDashboard';

import '../sass/main.scss';

const store = configureStore();
const history = syncHistoryWithStore(hashHistory, store);

render(
  <Provider store={store}>
    <MuiThemeProvider>
      <Router history={history}>
        <Route path="/" component={RootContainer}>
          <IndexRoute component={Introduction} />

          <Route path="/app-home" component={DecentralizedEqutyMarketHome} />
          <Route path="/human-authentication" component={HumanAuthentication} />
          <Route path="/user-home" component={ShipperHomeDashboard} />
          <Route path="/corporate-home" component={CorporateDashboard} />
          <Route path="/market-home" component={CarrierHomeDashboard} />

          <Route path="/market-share" component={MarketSahreBoard} />
          <Route path="/market-transaction" component={CarrierHomeDashboard} />
          <Route path="/market-bitask" component={CarrierHomeDashboard} />
        </Route>
      </Router>
    </MuiThemeProvider>
  </Provider>,
  document.getElementById('root')
);
