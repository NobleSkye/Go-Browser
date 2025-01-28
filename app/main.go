package main

import (
	"strings"

	webview "github.com/webview/webview_go"
)

func main() {
	// Create new webview instance
	w := webview.New(true)
	defer w.Destroy()

	// Set window properties
	w.SetTitle("Skye's Browser")
	w.SetSize(1024, 768, webview.HintNone)

	// Bind navigation function
	w.Bind("navigate", func(url string) {
		// Add http:// if no protocol is specified
		if !strings.HasPrefix(url, "http://") && !strings.HasPrefix(url, "https://") {
			url = "https://" + url
		}
		w.Navigate(url)
	})

	// Bind back and forward navigation functions
	w.Bind("goBack", func() {
		w.Eval("window.history.back();")
	})

	w.Bind("goForward", func() {
		w.Eval("window.history.forward();")
	})

	// Initial navigation bar setup - this runs before any page loads
	w.Init(`
        // Function to inject our nav bar
        function createNavBar() {
            // Check if nav already exists
            if (document.getElementById('skyeNavBar')) {
                return;
            }

            // Create container
            const nav = document.createElement('div');
            nav.id = 'skyeNavBar';
            nav.style.cssText = 'all: initial; position: fixed; top: 0; left: 0; right: 0; ' +
                               'height: 40px; background: #2b2b2b; border-bottom: 1px solid #404040; ' +
                               'display: flex; align-items: center; padding: 0 10px; z-index: 2147483647; ' +
                               'font-family: Arial, sans-serif;';

            // Create back button
            const backButton = document.createElement('button');
            backButton.textContent = 'Back';
            backButton.style.cssText = 'padding: 5px 15px; background: #404040; color: #fff; ' +
                                      'border: none; border-radius: 4px; cursor: pointer; ' +
                                      'font-size: 14px;';
            backButton.addEventListener('click', function() {
                goBack();
            });

            // Create forward button
            const forwardButton = document.createElement('button');
            forwardButton.textContent = 'Forward';
            forwardButton.style.cssText = 'padding: 5px 15px; background: #404040; color: #fff; ' +
                                         'border: none; border-radius: 4px; cursor: pointer; ' +
                                         'font-size: 14px;';
            forwardButton.addEventListener('click', function() {
                goForward();
            });

            // Create URL input
            const urlInput = document.createElement('input');
            urlInput.id = 'skyeUrlBar';
            urlInput.type = 'text';
            urlInput.style.cssText = 'flex: 1; margin: 0 10px; padding: 5px 10px; ' +
                                   'border: 1px solid #404040; border-radius: 4px; ' +
                                   'background: #3b3b3b; color: #fff; font-size: 14px; ' +
                                   'outline: none; min-width: 0;';
            urlInput.value = window.location.href;
            urlInput.addEventListener('keydown', function(e) {
                // Allow navigating when "Enter" is pressed
                if (e.key === 'Enter') {
                    let url = this.value;
                    if (!url.startsWith('http://') && !url.startsWith('https://')) {
                        url = 'https://' + url;
                    }
                    navigate(url);
                }
            });

            // Create navigation button
            const goButton = document.createElement('button');
            goButton.textContent = 'Search';
            goButton.style.cssText = 'padding: 5px 15px; background: #404040; color: #fff; ' +
                                   'border: none; border-radius: 4px; cursor: pointer; ' +
                                   'font-size: 14px;';
            goButton.addEventListener('click', function() {
                let url = urlInput.value;
                if (!url.startsWith('http://') && !url.startsWith('https://')) {
                    url = 'https://' + url;
                }
                navigate(url);
            });

            // Assemble the nav bar
            nav.appendChild(backButton);
            nav.appendChild(forwardButton);
            nav.appendChild(urlInput);
            nav.appendChild(goButton);

            // Add to page
            document.documentElement.appendChild(nav);

            // Add margin to body so the content is slightly below the nav bar
            if (document.body) {
                document.body.style.marginTop = '150px'; // Move content down by 150px
                document.body.style.paddingTop = '0';   // Reset padding-top to 0
                document.body.style.height = 'calc(100% - 150px)'; // Adjust content height accordingly
            }

        }

        // Create initial nav bar
        createNavBar();

        // Update URL in nav bar when it changes
        function updateUrl() {
            const urlInput = document.getElementById('skyeUrlBar');
            let currentUrl = window.location.href;

            // Visually change "google.com/search" to "skyesearch.cc/search"
            if (currentUrl.includes('google.com/search')) {
                currentUrl = currentUrl.replace('google.com/search', 'skyesearch.cc/search');
            }

        }

        // Observe DOM changes to ensure nav bar stays present
        const observer = new MutationObserver((mutations) => {
            createNavBar();
            updateUrl();
        });

        // Start observing document
        observer.observe(document.documentElement, {
            childList: true,
            subtree: true
        });

        // Backup interval for URL updates
        setInterval(() => {
            createNavBar();
            updateUrl();
        }, 1000);
    `)

	// Load the initial page
	w.Navigate("https://newtab.skyesearch.cc")

	// Run the webview
	w.Run()
}
