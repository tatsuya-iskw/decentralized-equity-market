// Data
import React, {Component} from 'react';
import {connect} from 'react-redux';
import {Link} from 'react-router';
// UI
import Grid from '@material-ui/core/Grid';
import Button from '@material-ui/core/Button';


function mapStateToProps(state) {
  return {};
}
const mapDispatchToProps = (dispatch, ownProps) => ({});

class Introduction extends Component {
  constructor(props) {
    super(props)
    this.state = {}
  }

  componentDidMount() {}

  render() {
    return (
      <div className="scdemo">
        <div style={{textAlign: "center", paddingTop: 128}}>
          <h1 style={{fontSize: 64}}>Decentralized Equity Market</h1>
          <h2><img src="https://www.ki-decentralized.com/img/ki-decentralized-logo.png" /></h2>
          <h2 style={{paddingTop: 48}}>
            <Link to="/app-home">
              <Button>Grasp a tny pieace of the future in Asset Swapping Market</Button>
            </Link>
          </h2>
        </div>
      </div>
    );
  }
}

export default connect(
  mapStateToProps,
  mapDispatchToProps
)(Introduction);