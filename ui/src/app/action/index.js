import fetch from 'isomorphic-fetch';

export const FETCH_PORTFOLIO = 'FETCH_PORTFOLIO';
export const fetchPortfolio = () => dispatch => {
  fetch('http://localhost:9998/api/portfolio/get')
  .then(response => response.json())
  .then(json => dispatch({
    type: FETCH_PORTFOLIO,
    data: json
  }));
};

export const FETCH_CORPORATE_SHARE = 'FETCH_CORPORATE_SHARE';
export const fetchCorporateShares = () => dispatch => {
  fetch('http://localhost:9998/api/share/get/cccccccccccccccccccccccccccc')
  .then(response => response.json())
  .then(json => dispatch({
    type: FETCH_CORPORATE_SHARE,
    data: json
  }));
};

export const FETCH_SCAN_CORP_INFO = 'FETCH_SCAN_CORP_INFO';
export const scanCorporation = () => dispatch => {
  console.log("ACTION - FETCH_SCAN_CORP_INFO")
  fetch('http://localhost:9998/api/corporate/scan/')
  .then(response => response.json())
  .then(json => dispatch({
    type: FETCH_SCAN_CORP_INFO,
    data: json
  }));
};

export const FETCH_SHARE_BY_USER_ID = 'FETCH_SHARE_BY_USER_ID';
export const fetchShareByUserId = (target_user_id) => dispatch => {
  // fetch('http://localhost:9998/api/share/get/' + target_user_id)
  fetch('http://localhost:9998/api/share/get/')
  .then(response => response.json())
  .then(json => dispatch({
    type: FETCH_SHARE_BY_USER_ID,
    data: json
  }));
};

export const BUY_SHARES = 'BUY_SHARES';
export const buyShares = () => dispatch => {
  fetch('http://localhost:9998/api/share/buy/corporate-id/2/4F30880E84638F3B4F2DCFE3CC53022CCAE10D491689E0DAC13B85C9C3A8F0E7')
  .then(response => response.json())
  .then(json => dispatch({
    type: BUY_SHARES,
    data: json
  }));
};

export const FETCH_ALL_SHARE_BY_REGULATOR = 'FETCH_ALL_SHARE_BY_REGULATOR';
export const fetchAllSharesRegulator = () => dispatch => {
  fetch('http://localhost:9998/api/share/get/all/')
  .then(response => response.json())
  .then(json => dispatch({
    type: FETCH_ALL_SHARE_BY_REGULATOR,
    data: json
  }));
};
export const FETCH_ALL_TRANSACTION_BY_REGULATOR = 'FETCH_ALL_TRANSACTION_BY_REGULATOR';
export const fetchAllTransactionRegulator = () => dispatch => {
  fetch('http://localhost:9998/api/transaction/get/all/')
  .then(response => response.json())
  .then(json => dispatch({
    type: FETCH_ALL_TRANSACTION_BY_REGULATOR,
    data: json
  }));
};
export const FETCH_ALL_BA_BY_REGULATOR = 'FETCH_ALL_BA_BY_REGULATOR';
export const fetchAllBARegulator = () => dispatch => {
  fetch('http://localhost:9998/api/bidask/get/all/')
  .then(response => response.json())
  .then(json => dispatch({
    type: FETCH_ALL_BA_BY_REGULATOR,
    data: json
  }));
};

export const ASSIGN_CARRIER = 'ASSIGN_CARRIER';
export const assignCarrier = (target_shipment_id) => dispatch => {
  console.log(target_shipment_id)
  fetch('http://localhost:9998/api/shipment/assigncarrier/' + target_shipment_id)
  .then(response => response.json())
  .then(json => dispatch({
    type: CREATE_SHIPMENT_ACTION,
    data: json
  }));
};

export const AUTHORIZE_TRANSACTION = 'AUTHORIZE_TRANSACTION';
export const authorizeTransaction = (target_shipment_id, current_status) => dispatch => {
  fetch('http://localhost:9998/api/shipment/status/update/' + target_shipment_id + '/' + current_status)
  .then(response => response.json())
  .then(json => dispatch({
    type: AUTHORIZE_TRANSACTION,
    data: json
  }));
};

export const UPDATE_SHIPMENT_STATUS_ON_SHIPPER = 'UPDATE_SHIPMENT_STATUS';
export const updateShipmentStatusOnShipper = (target_shipment_id, current_status) => dispatch => {
  fetch('http://localhost:9998/api/shipment/status/update/' + target_shipment_id + '/' + current_status)
  .then(response => response.json())
  .then(json => dispatch({
    type: UPDATE_SHIPMENT_STATUS_ON_SHIPPER,
    data: json
  }));
};

export const UPDATE_SHIPMENT_STATUS = 'UPDATE_SHIPMENT_STATUS';
export const updateShipmentStatusOnCarrier = (target_shipment_id, current_status) => dispatch => {
  console.log(target_shipment_id)
  console.log(current_status)

  fetch('http://localhost:9998/api/shipment/status/update/' + target_shipment_id + '/' + current_status)
  .then(response => response.json())
  .then(json => dispatch({
    type: UPDATE_SHIPMENT_STATUS,
    data: json
  }));
};

export const MAKE_PAYMENT = 'MAKE_PAYMENT';
export const makePayment = (target_shipment_id, current_shipment_status) => dispatch => {
  fetch('http://localhost:9998/api/shipment/status/update/' + target_shipment_id + '/' + current_shipment_status)
  .then(response => response.json())
  .then(json => dispatch({
    type: MAKE_PAYMENT,
    data: json
  }));
};

export const CREATE_SHIPMENT_ACTION = 'CREATE_SHIPMENT_ACTION';
export const createShipment = () => dispatch => {
  fetch('http://localhost:9998/api/shipment/create')
  .then(response => response.json())
  .then(json => dispatch({
    type: CREATE_SHIPMENT_ACTION,
    data: json
  }));
};

export const RECIEVE_INITIAL_DATA = 'RECIEVE_INITIAL_DATA';
export const fetchInitialData = () => dispatch => {
  fetch('http://localhost:9998/api/list/shipment')
  .then(response => response.json())
  .then(json => dispatch({
    type: RECIEVE_INITIAL_DATA,
    data: json
  }));
};

export const FETCH_ALL_SHIPMENT_LIST_ON_SHIPPER = 'FETCH_ALL_SHIPMENT_LIST_ON_SHIPPER';
export const fetchAllShipmentListOnShipper = (shipper_id) => dispatch => {
  console.log()
  fetch('http://localhost:9998/api/shipment/list/shipper/' + shipper_id)
  .then(response => response.json())
  .then(json => dispatch({
    type: FETCH_ALL_SHIPMENT_LIST_ON_SHIPPER,
    data: json
  }));
};

export const FETCH_ALL_SHIPMENT_LIST_ON_CARRIER = 'FETCH_ALL_SHIPMENT_LIST_ON_CARRIER';
export const fetchAllShipmentListOnCarrier = () => dispatch => {
  fetch('http://localhost:9998/api/shipment/list/carrier')
  .then(response => response.json())
  .then(json => dispatch({
    type: FETCH_ALL_SHIPMENT_LIST_ON_CARRIER,
    data: json
  }));
};

export const GET_ALL_HISTORICAL_LOCATIONS = 'GET_ALL_HISTORICAL_LOCATIONS';
export const getHistoricalLocations = (shipment_key) => dispatch => {
  console.log(shipment_key)
  console.log('http://localhost:9998/api/location/get/' + shipment_key)

  fetch('http://localhost:9998/api/location/get/' + shipment_key)
  .then(response => response.json())
  .then(json => dispatch({
    type: GET_ALL_HISTORICAL_LOCATIONS,
    data: json
  }));
};

export const LOG_SHIPMENT_LOCATION = 'LOG_SHIPMENT_LOCATION';
export const logShipmentLocation = (shipment_id) => dispatch => {
  fetch('http://localhost:9998/api/location/log/' + shipment_id)
  .then(response => response.json())
  .then(json => dispatch({
    type: LOG_SHIPMENT_LOCATION,
    data: json
  }));
};