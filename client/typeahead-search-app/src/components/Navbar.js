import { Navbar, Nav, Container, NavDropdown } from 'react-bootstrap';

function NavbarTop() {
    return (
    <Navbar bg="light" expand="lg">
    <Container>
      <Navbar.Brand href="/">FLASH</Navbar.Brand>
      <Navbar.Toggle aria-controls="basic-navbar-nav" />
      <Navbar.Collapse id="basic-navbar-nav">
        <Nav className="ml-auto">
          <Nav.Link href="/">Home</Nav.Link>
        </Nav>
        <Nav class="ms-auto">
          <Nav.Link href="/about">About</Nav.Link >
        </Nav>
      </Navbar.Collapse>
    </Container>
  </Navbar>
    )
}


export default NavbarTop;
