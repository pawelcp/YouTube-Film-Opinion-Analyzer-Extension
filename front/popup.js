document.getElementById('getOpinionButton').addEventListener('click', async () => {
    const [tab] = await chrome.tabs.query({ active: true, currentWindow: true });
    const videoLink = encodeURIComponent(tab.url);
    const apiUrl = `http://localhost:8080/get_opinion?link=${videoLink}`;
    console.log(videoLink)

    try {
        const response = await fetch(apiUrl, {
            mode: 'cors' // Use 'cors' mode
        });

        if (response.ok) {
            const data = await response.json();
            const opinionResult = document.getElementById('opinionResult');
            opinionResult.textContent = 'Opinion: ' + data.opinion;
        } else {
            throw new Error('Failed to fetch: ' + response.statusText);
        }
    } catch (error) {
        const opinionResult = document.getElementById('opinionResult');
        opinionResult.textContent = 'Error fetching opinion.';
        console.error('Error fetching opinion:', error);
    }
});
