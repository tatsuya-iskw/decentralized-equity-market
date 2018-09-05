import {createStore, applyMiddleware} from 'redux';
import thunkMiddleware from 'redux-thunk';
import promiseMiddleware from 'redux-promise';
import rootReducer from './reducers';

export default function configureStore() {
  const initialState = {
    corporateVisualization: {
      shares: []
    },
    shipperReducer: {
      portfolio: [],
      shares: [],
      corp_total: {},
      corp_own: {},
      shipmentList: []
    },
    shipmentListCarrier: {
      shares: [],
      bidasks: [],
      transactions: [],
      shipmentList: []
    }
  };
  
  return createStore(
    rootReducer,
    initialState,
    applyMiddleware(
      thunkMiddleware,
      promiseMiddleware
    )
  );
}
