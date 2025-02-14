import { useEffect, useState } from "react";
import PageHeader from "../header/PageHeader";
import axios from 'axios';

function BankList() {
    const [bankAccounts, setBankAccounts] = useState([{ id: '', holder_name: '', phone_no: '', account_type: '' }]);

    const readAllBankAccounts = async () => {
        try {
            const baseUrl = 'http://localhost:8080';
            const response = await axios.get(`${baseUrl}/bank`);
            const queriedBankAccounts = response.data;
            setBankAccounts(queriedBankAccounts);
        } catch (error) {
            alert('Server Error');
        }
    };

    const deleteBankAccount = async (id) => {
        if (!confirm("Are you sure to delete?")) {
            return;
        }
        const baseUrl = "http://localhost:8080";
        try {
            const response = await axios.delete(`${baseUrl}/bank/${id}`);
            alert(response.data.message);
            await readAllBankAccounts();
        } catch (error) {
            alert('Server Error');
        }
    };

    useEffect(() => {
        readAllBankAccounts();
    }, []);

    return (
        <>
            <PageHeader />
            <h3>List of Bank Accounts</h3>
            <div className="container">
                <table className="table table-success table-striped table-bordered">
                    <thead className="table-dark">
                        <tr>
                        <link href="https://cdn.jsdelivr.net/npm/bootstrap@5.3.3/dist/css/bootstrap.min.css" rel="stylesheet" integrity="sha384-QWTKZyjpPEjISv5WaRU9OFeRpok6YctnYmDr5pNlyT2bRjXh0JMhjY6hW+ALEwIH" crossorigin="anonymous"></link>
                        
                            <th scope="col">Account Number</th>
                            <th scope="col">Holder Name</th>
                            <th scope="col">Phone Number</th>
                            <th scope="col">Account Type</th>
                            <th></th>
                        </tr>
                    </thead>
                    <script src="https://cdn.jsdelivr.net/npm/bootstrap@5.3.3/dist/js/bootstrap.bundle.min.js" integrity="sha384-YvpcrYf0tY3lHB60NNkmXc5s9fDVZLESaAA55NDzOxhy9GkcIdslK1eN7N6jIeHz" crossorigin="anonymous"></script>
                    <tbody>
                        {(bankAccounts && bankAccounts.length > 0) ? bankAccounts.map(
                            (account) => {
                                return (
                                    <tr key={account.id}>
                                        <th scope="row">{account.id}</th>
                                        <td>{account.holder_name}</td>
                                        <td>{account.phone_no}</td>
                                        <td>{account.account_type}</td>
                                        <td>
                                            <a href={`/bank/view/${account.id}`} className="btn btn-success">View</a>
                                            &nbsp;
                                            <a href={`/bank/edit/${account.id}`} className="btn btn-warning">Edit</a>
                                            &nbsp;
                                            <button className="btn btn-danger" onClick={() => deleteBankAccount(account.id)}>Delete</button>
                                        </td>
                                    </tr>
                                );
                            }
                        ) : <tr><td colSpan="5">No Data Found</td></tr>}
                    </tbody>
                </table>
            </div>
        </>
    );
}

export default BankList;
