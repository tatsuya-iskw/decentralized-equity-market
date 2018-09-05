import {
  RECIEVE_INITIAL_DATA,
  CREATE_SHIPMENT_ACTION,
  FETCH_ALL_SHIPMENT_LIST_ON_CARRIER,
  UPDATE_SHIPMENT_STATUS,
  FETCH_ALL_SHIPMENT_LIST_ON_SHIPPER,
  UPDATE_SHIPMENT_STATUS_ON_SHIPPER,
  LOG_SHIPMENT_LOCATION,
  AUTHORIZE_TRANSACTION,
  MAKE_PAYMENT,
  GET_ALL_HISTORICAL_LOCATIONS,
  FETCH_SHARE_BY_USER_ID,
  FETCH_ALL_SHARE_BY_REGULATOR,
  FETCH_ALL_TRANSACTION_BY_REGULATOR,
  FETCH_ALL_BA_BY_REGULATOR,
  FETCH_SCAN_CORP_INFO,
  FETCH_CORPORATE_SHARE,
  FETCH_PORTFOLIO
} from '../action';

// import {combineReducers} from 'redux-immutable';
// https://github.com/gajus/redux-immutable
import {combineReducers} from 'redux';
// TODO: figure out why it gets error if it uses import
// import routerReducer from 'react-router-redux';
// using require is fine for 'react-router-redux'
const {routerReducer} = require('react-router-redux');

function shipperReducer(state = {}, action) {
  switch (action.type) {
    case FETCH_PORTFOLIO:
      return {
        portfolio: action.data,
        corp_total: state.corp_total,
        corp_own: state.corp_own,
        shares: state.shares
      };
    case FETCH_SCAN_CORP_INFO:
      return {
        portfolio: state.portfolio,
        corp_total: action.data[0],
        corp_own: action.data[1],
        shares: state.shares
      };
    case FETCH_SHARE_BY_USER_ID:
      return {
        portfolio: state.portfolio,
        corp_total: state.corp_total,
        corp_own: state.corp_own,
        shares: action.data
      };
    case CREATE_SHIPMENT_ACTION:
    case UPDATE_SHIPMENT_STATUS_ON_SHIPPER:
      return state
    case FETCH_ALL_SHIPMENT_LIST_ON_SHIPPER:
      return {
        shipmentList: action.data
      };
    case CREATE_SHIPMENT_ACTION:
      return {
        shipmentList: action.data
      };
    default:
      return state
  }
}

function shipmentDetail(state = {}, action) {
  switch (action.type) {
    case GET_ALL_HISTORICAL_LOCATIONS:
      return {
        locationList: action.data
      };
    default:
      return state
  }
}

function shipmentListRegulator(state = {}, action) {
  switch (action.type) {
    case AUTHORIZE_TRANSACTION:
      return state
    case RECIEVE_INITIAL_DATA:
      return {
        shipmentList: action.data
      };
    default:
      return state
  }
}

function corporateVisualization(state = {}, action) {
  switch (action.type) {
    case FETCH_CORPORATE_SHARE:
      return {
        shares: action.data
      }
    default:
      return state
  }
}

function shipmentListCarrier(state = {}, action) {
  switch (action.type) {
    case FETCH_ALL_SHARE_BY_REGULATOR:
      return {
        shares: action.data,
        bidasks: state.bidasks,
        transactions: state.transactions
      }
    case FETCH_ALL_TRANSACTION_BY_REGULATOR:
      return {
        shares: state.shares,
        bidasks: state.bidasks,
        transactions: action.data
      }
    case FETCH_ALL_BA_BY_REGULATOR:
      return {
        shares: state.shares,
        bidasks: action.data,
        transactions: state.transactions 
      }
    case LOG_SHIPMENT_LOCATION:
      return state
    case FETCH_ALL_SHIPMENT_LIST_ON_CARRIER:
      return {
        shipmentList: action.data
      };
    case UPDATE_SHIPMENT_STATUS:
      return state;
    default:
      return state
  }
}


const rootReducer = combineReducers({
  corporateVisualization,
  shipmentDetail,
  shipperReducer,
  shipmentListRegulator,
  shipmentListCarrier,
  routing: routerReducer
});

export default rootReducer;
