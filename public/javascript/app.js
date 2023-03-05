(() => {
	const answerWrapper = document.querySelectorAll('.answer-wrapper');
	const toggleBtns = document.querySelectorAll('.answer-toggle');

	for (const ans of answerWrapper) {
		ans.style.display = 'none';
	}

	for (const btn of toggleBtns) {
		btn.addEventListener('click', (e) => {
			const answer = e.target.parentElement.nextElementSibling;
			answer.style.display =
				answer.style.display === 'none' ? 'block' : 'none';
		});
	}
})();

// Get the editForm element and the current fact ID to be edited
const editForm = document.querySelector('#form-update-fact');
const factToEdit = editform && editForm.dataset.factid;

// Add an event listener to listen for the form submit
editForm && editForm.addEventListener('submit', (event) => {
    // Prevent the default behavior of the form element
    event.preventDefault();

    // Convert form data into a JS object
    const formData = Object.fromEntries(new FormData(editForm));

    return fetch(`/fact/${factToEdit}`, {
        // Use the PATCH method
        method: 'PATCH',
        headers: {
            'Content-Type': 'application/json',
        },
        // Convert the form's Object data into JSON
        body: JSON.stringify(formData),
    }).then(() => (document.location.href = `/fact/${factToEdit}`));
});

// Get deleteButton element and the current fact ID to delete
const deleteButton = document.querySelector('#delete-button')
const factToDelete = deleteButton && deleteButton.dataset.factid

// Add event listener to listen for button click
deleteButton && deleteButton.addEventListener('click', () => {
	// We ask the user if they are sure they want to delete the fact
    const result = confirm("Are you sure you want to delete this fact?")
    
	// If the user cancels the prompt, we exit here
    if (!result) return

    // If the user confirms that they want to delete, we send a DELETE request
    // URL uses the current fact's ID
    // Lastly, we redirect to index
    return fetch(`/fact/${factToDelete}`, { method: 'DELETE' })
            .then(() => document.location.href="/")
})

