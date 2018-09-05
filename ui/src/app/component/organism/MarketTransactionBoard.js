// DATA
import React, {Component} from 'react';
import {connect} from 'react-redux';
import {fetchAllSharesRegulator} from '../../action';
// UI
import Table from '@material-ui/core/Table';
import TableBody from '@material-ui/core/TableBody';
import TableCell from '@material-ui/core/TableCell';
import TableHead from '@material-ui/core/TableHead';
import TableRow from '@material-ui/core/TableRow';

function mapStateToProps(state) {
  return {
    shares: state.shipmentlistcarrier.shares
  };
}
const mapDispatchToProps = (dispatch, ownProps) => ({
  fetchAllSharesRegulator: () => dispatch(fetchAllSharesRegulator())
});

class CarrierHomeDashboard extends Component {
  constructor(props) {
    super(props)
    this.state = {}
  }

  componentDidMount() {
    this.props.fetchAllSharesRegulator();
  }

  render() {
    const sharesListArray = this.props.shares
    let shareCells = ""

    if (sharesListArray) {
      shareCells = sharesListArray.map((value, key) => {
        return (
          <TableRow>
            <TableCell>{value.Key}</TableCell>
            <TableCell>{value.Record.share_corporate_id}</TableCell>
            <TableCell>{value.Record.share_hash_id}</TableCell>
            <TableCell>{value.Record.share_owner_id}</TableCell>
          </TableRow>
        )
      })
    }

    return (
      <div classname="scdemo">
        <h1>Market Dashboard - All Shaers</h1>
        <Table>
          <TableHead>
            <TableRow>
              <TableCell>ledger key</TableCell>
              <TableCell>Corporation identity</TableCell>
              <TableCell>Share Identity</TableCell>
              <TableCell>Owner Identity</TableCell>
            </TableRow>
          </TableHead>
          <TableBody>
            {shareCells}
          </TableBody>
        </Table>
      </div>
    );
  }
}

export default connect(
  mapStateToProps,
  mapDispatchToProps
)(CarrierHomeDashboard);