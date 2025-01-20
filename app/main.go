package main

import (
    "github.com/webview/webview_go"
    "strings"
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
                if (e.key === 'Enter') {
                    navigate(this.value);
                }
            });

            // Create navigation button
            const goButton = document.createElement('button');
            goButton.textContent = 'Go';
            goButton.style.cssText = 'padding: 5px 15px; background: #404040; color: #fff; ' +
                                   'border: none; border-radius: 4px; cursor: pointer; ' +
                                   'font-size: 14px;';
            goButton.addEventListener('click', function() {
                navigate(urlInput.value);
            });

            // Assemble the nav bar
            nav.appendChild(urlInput);
            nav.appendChild(goButton);

            // Add to page
            document.documentElement.appendChild(nav);

            // Add margin to body
            if (document.body) {
                document.body.style.marginTop = '40px';
                document.body.style.paddingTop = '0';
                document.body.style.height = 'calc(100% - 40px)';
            }
        }

        // Create initial nav bar
        createNavBar();

        // Update URL in nav bar when it changes
        function updateUrl() {
            const urlInput = document.getElementById('skyeUrlBar');
            if (urlInput && urlInput.value !== window.location.href) {
                urlInput.value = window.location.href;
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
        }, 100);
    `)

    // Load the initial page
    w.Navigate("https://search.nobleskye.dev")

    // Run the webview
    w.Run()
}
