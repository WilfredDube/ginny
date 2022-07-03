const form = document.querySelector("form");

form.addEventListener("submit", async (event) => {
  event.preventDefault();

  const formData = new FormData(event.target);
  const formProps = Object.fromEntries(formData);

  console.log(formProps);

  axios
    .post(location.href, formProps)
    .then((response) => {
      window.location = response.request.responseURL;
    })
    .catch((error) => console.error(error));
});
