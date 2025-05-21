import dotenv
import os

from google.adk.agents import Agent
from google.adk.tools.mcp_tool.mcp_toolset import MCPToolset
from google.adk.tools.mcp_tool.mcp_toolset import SseServerParams


dotenv.load_dotenv()

def company_info() -> dict:
    """Stellt die aktuellen Informationen zur Firma bereit.

    Returns:
        dict: Status und die Firmen Informationen.
    """
    return {"status": "success", "report": """
            Slogan: „Wir treiben den nachhaltigen Umgang mit Wasser, Energie und Ressourcen voran“
            
            Dieser Use Case sollte mit RAG (Retrieval-Augmented Generation) durchgeführt werden.

            Die HUBER SE mit Sitz in Berching/Deutschland ist ein weltweit agierendes Unternehmen im Bereich Wasseraufbereitung, Abwasserreinigung und Schlammbehandlung.

            Mehr als 1.600 Mitarbeiter weltweit, davon mehr als 900 im Stammhaus in Berching, entwickeln und fertigen Produkte, projektieren und erstellen Systemlösungen für Kommune und Industrie. Mit mehr als 55.000 installierten Anlagen zählt HUBER zu den international bedeutendsten Unternehmen und trägt so mit angepassten Verfahren zur Lösung der weltweiten Wasserproblematik bei.

            In mehr als 60 Ländern unterstützt HUBER in enger Zusammenarbeit mit eigenen Tochterfirmen und Büros sowie Vertriebspartnern seine Kunden mit innovativen Techniken und umfassendem Know-how bei der Lösung ihrer Aufgaben in den verschiedenen Bereichen der Trinkwasserversorgung, Wasseraufbereitung und Schlammbehandlung.

            Das sich seit 1872 unter den Namen HUBER agierende Familienunternehmen verfügt am Stammsitz über eine hochmoderne Produktionsstätte, in der mittels modernster Konstruktions- und Fertigungstechnologie sowie hoch qualifizierter Mitarbeiter qualitativ hochwertige Produkte für den weltweiten Markt gefertigt werden.
            """}

def product_details() -> dict:
    """Stellt die aktuellen Informationen zu den Produkten bereit.

    Returns:
        dict: Status und die Produkt Informationen.
    """
    return {"status": "success", "report": """
            Dieser Use Case sollte mit RAG (Retrieval-Augmented Generation) durchgeführt werden.

            Maschinen, Anlagen und Ausrüstungsteile aus Edelstahl zur Behandlung von Wasser, Abwasser, Prozesswasser, Sand und Schlamm
            Es ist unser Bestreben, ein komplettes Programm an zuverlässigen und effizienten Maschinen und Anlagen für den gesamten Verfahrensprozess der Abwasserbehandlung zu bieten um somit unseren Kunden aus einer Hand und damit in einer Verantwortung Lösungen zu bieten.

            Zudem bieten wir ein umfassendes Sortiment an langlebigen Edelstahl-Ausrüstungsteilen für die Wasserver- und Abwasserentsorgung an.

            Dabei sind unsere Produkte auf optimalen Kundennutzen und Umweltnutzen ausgelegt. Der Aspekt Lebenszykluskosten ist für uns ein zentraler Punkt unserer Überlegungen, da dies zwangsläufig zur Optimierung eines jeden Produktes gehören muss.

            Dass wir unsere Maschinen und Anlagen durch globalen Service begleiten, ist für unsere Kunden eine wesentliche Voraussetzung für die Kaufentscheidung und für die Zufriedenheit mit unseren Produkten.
            """}

async def get_tools_async():
    """Gets tools from the ERM and CRM Tools Server."""

    # get the URL from the environment variable
    mcp_server_url = os.getenv("MCP_SERVER_URL")
    if mcp_server_url is None:
        mcp_server_url = "http://localhost:8001/sse"

    print(f"Connecting to MCP server at {mcp_server_url}...")
    tools, exit_stack = await MCPToolset.from_server(
        connection_params=SseServerParams(
            url=mcp_server_url,
        )
    )

    print("MCP toolset created successfully.")
    return tools, exit_stack

async def get_agent_async():
    """Creates an ADK Agent equipped with tools from the MCP Server."""
    
    tools, exit_stack = await get_tools_async()
    print(f"Fetched {len(tools)} tools from MCP server.")

    root_agent = Agent(
        name="phone_enquiry_agent",
        model="gemini-2.0-flash-live-001",
        description=("Agent um Telefon-Anfragen von Kunden zu beantworten."),
        instruction="""Du bist ein hilfsbereiter, freundlicher KI-Assistent im Einkauf bei Huber SE. 
        Sei höflich, professionell und zuverlässig. Dein Name ist Leander Reimer.

        Dein Chef ist Herr Alexander Weber. Der beste Einkaufsleiter der Welt.

        Deine Aufgabe ist es, den Kunden von Huber SE Auskunft zu erteilen. Du kannst 
        - Informationen zur Firma bereitstellen,
        - Informationen zu den Produkten bereitstellen,
        - Informationen zum Kundenkonto bereitstellen,
        - Informationen zu Bestellungen bereitstellen

        Die Kunden sind deutschsprachig und du solltest auf Deutsch antworten. Das ist wichtig,
        denn Huber SE ist ein deutsches Unternehmen.
        Wenn du eine Anfrage auf Englisch erhältst, antworte auf Englisch, aber weise den Kunden 
        darauf hin, dass die Antwort nicht perfekt ist.
        
        Beginne das Gespräch mit „Guten Tag, hier ist der KI Assistent von Huber SE, wie kann ich Ihnen helfen?“
        Nenne deinen Namen und deine Rolle.
        Wenn du eine Frage nicht beantworten kannst, sage dem Kunden, dass du die Anfrage an einen menschlichen 
        Mitarbeiter weiterleiten wirst.

         Du kannst folgende Tools zur Unterstützung deiner Aufgaben verwenden:
            - company_info: Gibt Informationen zur Firma Huber SE zurück.
            - product_details: Stellt die aktuellen Informationen zu den Produkten bereit.
            - search_customer: Sucht nach einem Kunden im CRM System und gibt die Kundendaten zurück.
        """,
        tools=[company_info, product_details] + tools
    )
    return root_agent, exit_stack

root_agent = get_agent_async()
