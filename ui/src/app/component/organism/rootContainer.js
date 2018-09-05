// 'Container Components'
// http://redux.js.org/docs/basics/UsageWithReact.html
import React, {Component} from 'react';
import {connect} from 'react-redux';
import {Link} from 'react-router';

const mapStateToProps = (state, ownProps) => ({
  // state is an Immutable object
});

// The available actions for the component
const mapDispatchToProps = (dispatch, ownProps) => ({
});

class rootContainer extends Component {
  render() {
    return (
      <div className="rootContainer">
        <div className="rootContainer__sidebar">
          <Link to="/">
            <h2 className="rootContainer__sidebar__title">Decentralized Equity Market</h2>
          </Link>
        </div>
        <div className="rootContainer__main">
          {this.props.children}
        </div>
      </div>
    );
  }
}

export default connect(mapStateToProps, mapDispatchToProps)(rootContainer);
