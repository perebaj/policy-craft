"""create policies table

Revision ID: 9109c54286d9
Revises:
Create Date: 2024-03-11 22:26:16.765677

"""
from typing import Sequence, Union

from alembic import op
import sqlalchemy as sa


# revision identifiers, used by Alembic.
revision: str = '9109c54286d9'
down_revision: Union[str, None] = None
branch_labels: Union[str, Sequence[str], None] = None
depends_on: Union[str, Sequence[str], None] = None


def upgrade() -> None:
    # create policies table
    op.create_table(
        'policies',
        sa.Column('id', sa.Integer, primary_key=True),
        sa.Column('name', sa.String, nullable=False),
        sa.Column('criteria', sa.String, nullable=False),
        sa.Column('value', sa.Integer, nullable=False),
        sa.Column('created_at', sa.DateTime, server_default=sa.func.now()),
        sa.Column('updated_at', sa.DateTime, server_default=sa.func.now(), server_onupdate=sa.func.now()),
    )


def downgrade() -> None:
    op.drop_table('policies')
