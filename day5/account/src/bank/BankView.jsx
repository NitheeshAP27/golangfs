import { useEffect, useState } from "react";
import { useParams } from "react-router-dom";
import PageHeader from "../header/PageHeader";
import axios from 'axios';

function BankView() {
    const [bankAccount, setBankAccount] = useState({ id: '', holder_name: '', phone_no: '', account_type: '' });
    const params = useParams();

    const readById = async () => {
        const baseUrl = "http://localhost:8080";
        try {
            const response = await axios.get(`${baseUrl}/bank/${params.id}`);
            const queriedBankAccount = response.data;
            setBankAccount(queriedBankAccount);
        } catch (error) {
            alert('Server Error');
        }
    };

    useEffect(() => {
        readById();
    }, []);

    return (
        <>
            <PageHeader />
            <h3><a href="/bank/list" className="btn btn-light">Go Back</a>View Bank Account</h3>
            <div className="container">
                <div className="form-group mb-3">
                    <label htmlFor="holder_name" className="form-label">Holder Name:</label>
                    <div className="form-control" id="holder_name">{bankAccount.holder_name}</div>
                </div>
                <div className="form-group mb-3">
                    <label htmlFor="phone_no" className="form-label">Phone Number:</label>
                    <div className="form-control" id="phone_no">{bankAccount.phone_no}</div>
                </div>
                <div className="form-group mb-3">
                    <label htmlFor="account_type" className="form-label">Account Type:</label>
                    <div className="form-control" id="account_type">{bankAccount.account_type}</div>
                </div>
            </div>
        </>
    );
}

export default BankView;
