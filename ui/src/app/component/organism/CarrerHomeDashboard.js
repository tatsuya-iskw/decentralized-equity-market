// DATA
import React, {Component} from 'react';
import {connect} from 'react-redux';
import {fetchAllSharesRegulator} from '../../action';
import {fetchAllTransactionRegulator} from '../../action';
import {fetchAllBARegulator} from '../../action';

// UI
import Table from '@material-ui/core/Table';
import TableBody from '@material-ui/core/TableBody';
import TableCell from '@material-ui/core/TableCell';
import TableHead from '@material-ui/core/TableHead';
import TableRow from '@material-ui/core/TableRow';

function mapStateToProps(state) {
  return {
    shares: state.shipmentListCarrier.shares,
    transactions: state.shipmentListCarrier.transactions,
    bidasks: state.shipmentListCarrier.bidasks
  };
}
const mapDispatchToProps = (dispatch, ownProps) => ({
  fetchAllSharesRegulator: () => dispatch(fetchAllSharesRegulator()),
  fetchAllTransactionRegulator: () => dispatch(fetchAllTransactionRegulator()),
  fetchAllBARegulator: () => dispatch(fetchAllBARegulator())
});

const styles = theme => ({
  root: {
    width: '100%',
    marginTop: theme.spacing.unit * 3,
    overflowX: 'auto',
  },
  table: {
    minWidth: 700,
  },
});

class CarrierHomeDashboard extends Component {
  constructor(props) {
    super(props)
    this.state = {}
  }

  componentDidMount() {
    this.props.fetchAllSharesRegulator();
    this.props.fetchAllTransactionRegulator();
    this.props.fetchAllBARegulator();
  }

  render() {
    const sharesListArray = this.props.shares
    const transactionsListArray = this.props.transactions
    const bidasksListArray = this.props.bidasks

    let shareCells = ""
    let transactionCells = ""
    let bidaskCells = ""

    if (sharesListArray) {

      shareCells = sharesListArray.map((value, key) => {
        let styles = {
          corpColor: {},
          ownerColor: {},
        }
        if (value.Record.share_corporate_id === "ab06a741577bf44e459b8c0c163196727f0e0ae1649db3e05c0cb5a3d6213def") {
          styles = {corpColor: {color: "#2E7D32"}}
        }

        if (value.Record.share_owner_id === "ab06a741577bf44e459b8c0c163196727f0e0ae1649db3e05c0cb5a3d6213def") {
          styles = {
            ownerColor: {color: "#2E7D32"},
            corpColor: styles.corpColor
          }
        } else if (value.Record.share_owner_id === "4F30880E84638F3B4F2DCFE3CC53022CCAE10D491689E0DAC13B85C9C3A8F0E7") {
          styles = {
            ownerColor: {color: "#1565C0"},
            corpColor: styles.corpColor
          }
        }

        return (
          <TableRow>
            <TableCell>{value.Key}</TableCell>
            <TableCell style={styles.corpColor}>{value.Record.share_corporate_id}</TableCell>
            <TableCell>{value.Record.share_hash_id}</TableCell>
            <TableCell style={styles.ownerColor}>{value.Record.share_owner_id}</TableCell>
          </TableRow>
        )
      })
    }

    if (transactionsListArray) {
      transactionCells = transactionsListArray.map((value, key) => {
        return (
          <TableRow>
            <TableCell>{value.Key}</TableCell>
            <TableCell>{value.Record.txn_from_user_id}</TableCell>
            <TableCell>{value.Record.txn_hash_id}</TableCell>
            <TableCell>{value.Record.txn_number}</TableCell>
            <TableCell>{value.Record.txn_price}</TableCell>
            <TableCell>{value.Record.txn_timestamp_utc}</TableCell>
            <TableCell>{value.Record.txn_to_user_id}</TableCell>
          </TableRow>
        )
      })
    }

    if (bidasksListArray) {
      bidaskCells = bidasksListArray.map((value, key) => {
        return (
          <TableRow>
            <TableCell>{value.Key}</TableCell>
            <TableCell>{value.Record.ba_number}</TableCell>
            <TableCell>{value.Record.ba_price}</TableCell>
            <TableCell>{value.Record.ba_share_hash_id}</TableCell>
            <TableCell>{value.Record.ba_timestamp_utc}</TableCell>
            <TableCell>{value.Record.ba_type}</TableCell>
            <TableCell>{value.Record.ba_user_id}</TableCell>
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

        <h1>Market Dashboard - All Transactions</h1>
        <Table>
          <TableHead>
            <TableRow>
              <TableCell>ledger key</TableCell>
              <TableCell>from_user</TableCell>
              <TableCell>txn_id</TableCell>
              <TableCell>number</TableCell>
              <TableCell>price</TableCell>
              <TableCell>timestamp</TableCell>
              <TableCell>to_user</TableCell>
            </TableRow>
          </TableHead>
          <TableBody>
            {transactionCells}
          </TableBody>
        </Table>

        <h1>Market Dashboard - All Bid Asks</h1>
        <Table>
          <TableHead>
            <TableRow>
              <TableCell>ledger key</TableCell>
              <TableCell>Number</TableCell>
              <TableCell>price</TableCell>
              <TableCell>share_hash_id</TableCell>
              <TableCell>timestamp</TableCell>
              <TableCell>transaction type</TableCell>
              <TableCell>user_hash_id</TableCell>
            </TableRow>
          </TableHead>
          <TableBody>
            {bidaskCells}
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