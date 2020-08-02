# image-to-ascii
This program takes the provided png or jpeg file and outputs it on the command line as ASCII art.

## Usage
 ./ascii -image [file name] -width [int] -height [int]</br>
 > -image [string] </br>
        Required: the path of the image to turn into ASCII art. </br>
    -width [int] </br>
        The width of the resulting ASCII art. (default 80) </br>
   -height [int] </br>
        The height of the resulting ASCII art. (default 0) </br>
        
  If either width or height (but not both) is set to 0, the resulting ASCII art will be scaled to be in the original ratio.
  
  ## Image formats
  This tool only supports png and jpg images.

  ## Acknowledgements
  Thanks to Github user @nfnt for creating the resize package used in this tool.