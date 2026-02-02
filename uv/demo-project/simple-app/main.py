"""
Simple demo application for UV exercises.
"""

def greet(name: str = "World") -> str:
    """Return a greeting message."""
    return f"Hello, {name}!"


def main():
    """Main entry point."""
    print(greet())
    print(greet("UV User"))
    
    # Demonstrate imports
    try:
        import requests
        print(f"\n✅ Requests version: {requests.__version__}")
    except ImportError:
        print("\n❌ Requests not installed. Run: uv add requests")
    
    try:
        from rich import print as rprint
        rprint("\n✅ [bold green]Rich is installed![/bold green]")
    except ImportError:
        print("\n❌ Rich not installed. Run: uv add rich")


if __name__ == "__main__":
    main()
