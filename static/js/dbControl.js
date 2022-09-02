const apiUrl = "http://127.0.0.1:8080/costumers";

var FirstName;
var LastName;
var Email;
var PhoneNumber;

// post costumer function
const postCostumer = (FirstNamePost, LastNamePost, EmailPost, PhoneNumberPost) => {
    fetch(apiUrl, {
        method: "post",
        headers: {
            'Accept': 'application/json',
            'Content-Type': 'application/json'
        },
        body: JSON.stringify({
            FirstName: FirstNamePost,
            LastName: LastNamePost,
            Email: EmailPost,
            PhoneNumber: PhoneNumberPost
        })
    })
    .then((response)=>{response.json()});
};
// send costumer
const sendCostumer = () => {
    let inputFirstName   = document.getElementById("FName").value;
    let inputLastName    = document.getElementById("LName").value;
    let inputEmail       = document.getElementById("Email").value;
    let inputPhoneNumber = document.getElementById("PhoneNumber").value;
    let formPost         = document.getElementById("post-costumer");
    postCostumer(inputFirstName, inputLastName, inputPhoneNumber, inputEmail);
    formPost.reset();
};
// delete costumer
const deleteCostumer = () => {
    let inputId = document.getElementById("costumer-id-delete").value;
    let deleteIdForm = document.getElementById("delete-id-form");
    if(isNaN(inputId)){
        alert("Invalid Input, not a number");
    }else{
        fetch(apiUrl+`/${inputId}`, {
            method: "DELETE",
            headers: {
                'Accept': 'application/json',
                'Content-Type': 'application/json'
            }
        })
        .then(response => response.text())
        .then(response => console.log(response));
    };
    deleteIdForm.reset();
};
// Update costumer
const updateCostumer = () => {
    let costumerID        = document.getElementById("CostumerID").value;
    let updateFirstName   = document.getElementById("UFName").value;
    let updateLastName    = document.getElementById("ULName").value;
    let updateEmail       = document.getElementById("UpdateEmail").value;
    let updatePhoneNumber = document.getElementById("UPhoneNumber").value;
    let formUpdate        = document.getElementById("update-costumer");

    fetch(apiUrl+`/${costumerID}`, {
        method: 'PUT',
        headers: {
            'Accept': 'application/json',
            'Content-Type': 'application/json'
        },
        body: JSON.stringify({
            FirstName: updateFirstName,
            LastName: updateLastName,
            Email: updateEmail,
            PhoneNumber: updatePhoneNumber
        })
    });
    formUpdate.reset();
};