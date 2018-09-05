// Data
import React, {Component} from 'react';
import {connect} from 'react-redux';
import {Link} from 'react-router';
// UI
import Grid from '@material-ui/core/Grid';

import aliceIcon from '../img/alice.svg'
import bobIcon from '../img/bob.svg'
import regIcon from '../img/reg.svg'
import Card from '@material-ui/core/Card';
import CardActions from '@material-ui/core/CardActions';
import CardContent from '@material-ui/core/CardContent';

import Button from '@material-ui/core/Button';
function mapStateToProps(state) {
  return {
    visualizeData: state.visualizeData
  };
}
const mapDispatchToProps = (dispatch, ownProps) => ({
  fetchInitialData: () => dispatch(fetchInitialData())
});

class SupplyChainDemoScreen extends Component {
  constructor(props) {
    super(props)
    this.state = {}
  }

  componentDidMount() {}

  render() {
    return (
      <div className="scdemo">
        <div style={{}}>
          <div style={{margin: 36}}>
          <Card>
            <CardContent>
            <Link to="/human-authentication">
              <svg className="">
                <use xlinkHref={bobIcon} />
              </svg>
              <Button style={{fontSize: 44}}>Bob (Human)</Button>
            </Link>
              <h3 style={{color: "#1565C0"}}>digital Identity = 4F30880E84638F3B4F2DCFE3CC53022CCAE10D491689E0DAC13B85C9C3A8F0E7</h3>
            </CardContent>
          </Card>
          </div>
          <div style={{margin: 36}}>
          <Card>
            <CardContent>
              <Link to="/corporate-home">
                <svg className="">
                  <use xlinkHref={aliceIcon} />
                </svg>
                <Button style={{fontSize: 44}}>Alice (Public Corporation)</Button>
              </Link>
              <h3 style={{color: "#2E7D32"}}>digital Identity = ab06a741577bf44e459b8c0c163196727f0e0ae1649db3e05c0cb5a3d6213def</h3>
            </CardContent>
          </Card>
          </div>
          <div style={{margin: 36}}>
          <Card>
            <CardContent>
            <Link to="/market-home">
                <svg className="">
                  <use xlinkHref={regIcon} />
                </svg>
              <Button style={{fontSize: 44}}>Regulator</Button>
            </Link>
            </CardContent>
          </Card>
          </div>
        </div>
      </div>
    );
  }
}

export default connect(
  mapStateToProps,
  mapDispatchToProps
)(SupplyChainDemoScreen);