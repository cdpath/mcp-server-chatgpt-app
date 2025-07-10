import subprocess
import os
import tempfile
from typing import Optional, Dict, Any
from mcp.server.fastmcp import FastMCP

mcp = FastMCP("ChatGPT")


def _run_shortcuts_command(
    temp_file_path: str, wait_for_output: bool
) -> Dict[str, Any]:
    """
    Execute the shortcuts command with proper error handling.

    Args:
        temp_file_path: Path to the temporary file containing the prompt
        wait_for_output: Whether to wait for and capture output

    Returns:
        Dict containing operation status and response
    """
    command = ["shortcuts", "run", "Ask ChatGPT on Mac", "--input-path", temp_file_path]

    try:
        if wait_for_output:
            result = subprocess.run(command, capture_output=True, text=True, check=True)
            return {
                "operation": "ask_chatgpt",
                "status": "success",
                "message": result.stdout.strip(),
            }
        else:
            subprocess.Popen(
                command,
                stdout=subprocess.DEVNULL,
                stderr=subprocess.DEVNULL,
                start_new_session=True,
            )
            return {
                "operation": "ask_chatgpt",
                "status": "success",
                "message": "Sent to ChatGPT (not waiting for response)",
            }
    except subprocess.CalledProcessError as e:
        return {
            "operation": "ask_chatgpt",
            "status": "error",
            "message": f"Shortcuts command failed: {e.stderr.strip() if e.stderr else str(e)}",
        }
    except Exception as e:
        return {
            "operation": "ask_chatgpt",
            "status": "error",
            "message": f"Unexpected error: {str(e)}",
        }


@mcp.tool()
def ask_chatgpt(prompt: str, wait_for_output: bool = False) -> Dict[str, Any]:
    """
    Send a prompt to ChatGPT macOS app using Shortcuts.

    Args:
        prompt: The text to send to ChatGPT
        wait_for_output: Whether to wait for ChatGPT to respond
    Returns:
        Dict containing operation status and response
    """
    with tempfile.NamedTemporaryFile(
        mode="w", suffix=".txt", delete=False
    ) as temp_file:
        temp_file.write(prompt)
        temp_file.flush()
        temp_file_path = temp_file.name

    return _run_shortcuts_command(temp_file_path, wait_for_output=wait_for_output)


def main():
    """Entry point for the MCP server."""
    mcp.run()


if __name__ == "__main__":
    main()
