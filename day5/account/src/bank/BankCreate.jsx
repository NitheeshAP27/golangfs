import { useState } from "react";
import PageHeader from "../header/PageHeader";
import { useNavigate } from "react-router-dom";
import axios from 'axios';

function BankCreate() {
    const [bankDetails, setBankDetails] = useState({ id: '', holder_name: '', phone_no: '', account_type: '' });
    const navigate = useNavigate();

    const txtBoxOnChange = event => {
        const updatableBankDetails = { ...bankDetails };
        updatableBankDetails[event.target.id] = event.target.value;
        setBankDetails(updatableBankDetails);
    };

    const createBankDetails = async () => {
        const baseUrl = "http://localhost:8080";
        try {
            const response = await axios.post(`${baseUrl}/bank`, { ...bankDetails });
            const createdBankDetails = response.data.bank;
            setBankDetails(createdBankDetails);
            alert(response.data.message);
            navigate('/bank/list'); // Update this to the appropriate route for listing bank accounts
        } catch (error) {
            alert('Server Error');
        }
    };

    return (
        <>
        
            <PageHeader />
            <h3><a href="/bank/list" className="btn btn-light">Go Back</a> Add Bank Account</h3>
            <div className="container">
                <div className="form-group mb-3">
                    <label htmlFor="holder_name" className="form-label">Holder Name:</label>
                    <input type="text" className="form-control" id="holder_name" 
                        placeholder="Please enter holder name"
                        value={bankDetails.holder_name} 
                        onChange={txtBoxOnChange} />
                </div>
                <div className="form-group mb-3">
                    <label htmlFor="phone_no" className="form-label">Phone Number:</label>
                    <input type="text" className="form-control" id="phone_no" 
                        placeholder="Please enter phone number"
                        value={bankDetails.phone_no} 
                        onChange={txtBoxOnChange} />
                </div>
                <div className="form-group mb-3">
                    <label htmlFor="account_type" className="form-label">Account Type:</label>
                    <input type="text" className="form-control" id="account_type" 
                        placeholder="Please enter account type"
                        value={bankDetails.account_type} 
                        onChange={txtBoxOnChange} />
                </div>
                <button className="btn btn-primary"
                    onClick={createBankDetails}>Create Bank Account</button>
            </div>
        </>
    );
}

export default BankCreate;
