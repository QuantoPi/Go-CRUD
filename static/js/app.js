const getCostumers = async () => {
    let url = "http://127.0.0.1:8080/costumers";
    try {
        let response = await fetch(url);
        return await response.json();
    }
    catch (error) {
        console.log(error);
    }
}
const costumersTable = async ()=>{
    let costumers = await getCostumers();
    var apiTable = document.getElementById('api-table');
    let row = apiTable.insertRow(0);
    let IDCell = row.insertCell(0);
    let firstNameCell = row.insertCell(1);
    let lastNameCell = row.insertCell(2);
    let emailCell = row.insertCell(3);
    let phoneCell = row.insertCell(4);

    IDCell.innerHTML = "ID";
    firstNameCell.innerHTML = "First Name";
    lastNameCell.innerHTML = "Last Name";
    emailCell.innerHTML = "Email";
    phoneCell.innerHTML = "Phone Number";

    var iterator = 1 ;
    costumers.forEach(costumer => {
        let row = apiTable.insertRow(iterator);
        let IDCell = row.insertCell(0);
        let firstNameCell = row.insertCell(1);
        let lastNameCell = row.insertCell(2);
        let emailCell = row.insertCell(3);
        let phoneCell = row.insertCell(4);
    
        IDCell.innerHTML = costumer.ID;
        firstNameCell.innerHTML = costumer.FirstName;
        lastNameCell.innerHTML = costumer.LastName;
        emailCell.innerHTML = costumer.Email;
        phoneCell.innerHTML = costumer.PhoneNumber;  
        iterator+=1;
    });
}
costumersTable();