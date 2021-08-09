from jinja2 import Environment, FileSystemLoader
import os


file_loader = FileSystemLoader('utils/generators/templates')
env = Environment(loader=file_loader)


def get_input_values():
    """
    Get input values from user

    :return: cmp_name, cmp_folder, attributes
    """
    cmp_name = input("Enter component name: ")
    attributes = input("Enter attributes ( name:[type] ) : ")
    attributes = [{"name": x.split(':')[0], "type": x.split(':')[1]} for x in attributes.split(' ')]
    return cmp_name, attributes


def gen_api_templates(cmp_name, attributes):
    """
    Generate API template files

    :param cmp_name: Component name
    :param attributes: Attributes
    """
    mkdir('output')
    # Model
    gen_template(
        template_file='model.j2',
        output_file='models/'+cmp_name.lower()+'.go',
        cmp_name=cmp_name, attributes=attributes)
    # Repository
    gen_template(
        template_file='repository.j2',
        output_file='internal/repository/'+cmp_name.lower()+'-repository.go',
        cmp_name=cmp_name, attributes=attributes)
    # Controller
    gen_template(
        template_file='controller.j2',
        output_file='internal/controllers/'+cmp_name.lower()+'-controller.go',
        cmp_name=cmp_name, attributes=attributes)
    # Routes
    gen_template(
        template_file='routes.j2',
        output_file='internal/routes/'+cmp_name.lower()+'-routes.go',
        cmp_name=cmp_name, attributes=attributes)


def gen_template(template_file, output_file, cmp_name, attributes):
    """
    Generate a file from a template file and fill it with the
    attributes.

    :param template_file: Template file to use
    :param output_file: Output file to write to
    :param cmp_name: Component name
    :param attributes: Attributes to fill the template with
    """
    template = env.get_template(template_file)
    output = template.render(cmp_name=cmp_name, attributes=attributes)
    with open(output_file, 'w') as f:
        f.write(output)


def mkdir(routes_dir):
    """
    Create a directory if it does not exist.

    :param routes_dir: Directory to create
    """
    if not os.path.exists(routes_dir):
        os.makedirs(routes_dir)


if __name__ == '__main__':
    cmp_name, attributes = get_input_values()
    gen_api_templates(cmp_name, attributes)
    