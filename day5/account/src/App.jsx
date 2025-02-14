
import BankList from './bank/BankList'
import BankCreate from './bank/BankCreate'
import BankView from './bank/BankView'

import { BrowserRouter, Route, Routes } from 'react-router-dom'
import BankEdit from './bank/BankEdit';


function App() {
  return (
    <>
      <div>
        <BrowserRouter>
          <Routes>
            <Route path="" element={<BankList/>}/>
            <Route path="/bank/list" element={<BankList/>}/>
            <Route path="/bank/create" element={<BankCreate/>}/>
            <Route path="/bank/view/:id" element={<BankView/>}/>
            <Route path="/bank/edit/:id" element={<BankEdit/>}/>
          </Routes>
        </BrowserRouter>
      </div>
    </>
  );
}

export default App;
