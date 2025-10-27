import { useRef, useEffect } from "react";

export const useOutClick = (callback: () => void) => {
    const ref = useRef<HTMLDivElement>(null); // Create a ref to attach to your target div

    useEffect(() => {
        const handleClickOutside = (event: MouseEvent | TouchEvent) => {
        // Check if the ref exists and if the clicked element is NOT inside the ref's element
        if (ref.current && !ref.current.contains(event.target as Node)) {
            callback(); // Execute the provided callback function
        }
        };

        // Add event listeners to the document for both mouse and touch events
        document.addEventListener('mousedown', handleClickOutside);
        document.addEventListener('touchstart', handleClickOutside);

        // Clean up the event listeners when the component unmounts
        return () => {
        document.removeEventListener('mousedown', handleClickOutside);
        document.removeEventListener('touchstart', handleClickOutside);
        };
    }, [callback]); // Re-run effect if the callback changes

    return ref; // Return the ref to be attached to your div
};